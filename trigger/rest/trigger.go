package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/project-flogo/contrib/trigger/rest/cors"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

const (
	CorsPrefix = "REST_TRIGGER"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{}, &Reply{})
var hanlderSettings *HandlerSettings

func init() {
	trigger.Register(&Trigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{id: config.Id, settings: s}, nil
}

// Trigger REST trigger struct
type Trigger struct {
	server   *Server
	settings *Settings
	id       string
	logger   log.Logger
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()

	router := httprouter.New()

	addr := ":" + strconv.Itoa(t.settings.Port)

	pathMap := make(map[string]string)

	preflightHandler := &PreflightHandler{logger: t.logger, c: cors.New(CorsPrefix, t.logger)}

	// Init handlers
	for _, handler := range ctx.GetHandlers() {
		hanlderSettings = &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), hanlderSettings, true)
		if err != nil {
			return err
		}

		method := hanlderSettings.Method
		path := hanlderSettings.Path

		t.logger.Debugf("Registering handler [%s: %s]", method, path)

		if _, ok := pathMap[path]; !ok {
			pathMap[path] = path
			router.OPTIONS(path, preflightHandler.handleCorsPreflight) // for CORS
		}

		//router.OPTIONS(path, handleCorsPreflight) // for CORS
		router.Handle(method, path, newActionHandler(t, handler))
	}

	t.logger.Debugf("Configured on port %d", t.settings.Port)
	t.server = NewServer(addr, router)

	return nil
}

func (t *Trigger) Start() error {
	return t.server.Start()
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	return t.server.Stop()
}

type PreflightHandler struct {
	logger log.Logger
	c      cors.Cors
}

// Handles the cors preflight request
func (h *PreflightHandler) handleCorsPreflight(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	h.logger.Infof("Received [OPTIONS] request to CorsPreFlight: %+v", r)
	h.c.HandlePreflight(w, r)
}

// IDResponse id response object
type IDResponse struct {
	ID string `json:"id"`
}

func newActionHandler(rt *Trigger, handler trigger.Handler) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		rt.logger.Debugf("Received request for id '%s'", rt.id)

		c := cors.New(CorsPrefix, rt.logger)
		c.WriteCorsActualRequestHeaders(w)

		out := &Output{}

		out.PathParams = make(map[string]string)
		for _, param := range ps {
			out.PathParams[param.Key] = param.Value
		}

		queryValues := r.URL.Query()
		out.QueryParams = make(map[string]string, len(queryValues))
		out.Headers = make(map[string]string, len(r.Header))

		for key, value := range r.Header {
			out.Headers[key] = strings.Join(value, ",")
		}

		for key, value := range queryValues {
			out.QueryParams[key] = strings.Join(value, ",")
		}

		// Check the HTTP Header Content-Type
		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/x-www-form-urlencoded":
			buf := new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			s := buf.String()
			m, err := url.ParseQuery(s)
			content := make(map[string]interface{}, 0)
			if err != nil {
				rt.logger.Errorf("Error while parsing query string: %s", err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			for key, val := range m {
				if len(val) == 1 {
					content[key] = val[0]
				} else {
					content[key] = val[0]
				}
			}

			out.Content = content
		case "application/json":
			var content interface{}
			err := json.NewDecoder(r.Body).Decode(&content)
			if err != nil {
				switch {
				case err == io.EOF:
					// empty body
					//todo should handler say if content is expected?
				case err != nil:
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}
			out.Content = content
		default:

			if strings.Contains(contentType, "multipart/form-data") && hanlderSettings.File != "" {

				r.ParseMultipartForm(5 * 1024 * 1024)
				file, header, err := r.FormFile(hanlderSettings.File)

				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				result := make(map[*multipart.FileHeader]multipart.File)
				result[header] = file

				out.Content = result

			} else {
				b, err := ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				out.Content = string(b)
			}

		}

		results, err := handler.Handle(context.Background(), out)

		reply := &Reply{}
		reply.FromMap(results)

		if err != nil {
			rt.logger.Debugf("Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if reply.Data != nil {

			if reply.Code == 0 {
				reply.Code = 200
			}

			switch t := reply.Data.(type) {
			case string:
				var v interface{}
				err := json.Unmarshal([]byte(t), &v)
				if err != nil {
					//Not a json
					w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
				} else {
					//Json
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				}

				w.WriteHeader(reply.Code)
				_, err = w.Write([]byte(t))
				if err != nil {
					log.Error(err)
				}
				return
			default:
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(reply.Code)
				if err := json.NewEncoder(w).Encode(reply.Data); err != nil {
					log.Error(err)
				}
				return
			}
		}

		if reply.Code > 0 {
			w.WriteHeader(reply.Code)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
