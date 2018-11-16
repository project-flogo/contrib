package eftl

import (
	//"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"bytes"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"github.com/project-flogo/core/trigger"
	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/data/metadata"
	"github.com/mashling/commons/lib/eftl"
	"github.com/mashling/commons/lib/util"
)

const (
	settingURL      = "url"
	settingID       = "id"
	settingUser     = "user"
	settingPassword = "password"
	settingCA       = "ca"
	settingDest     = "dest"
	DefaultPort 	= 8181
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{}, &Reply{})
var logger log.Logger

func init() {
	trigger.Register(&Trigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// Trigger is a simple EFTL trigger
type Trigger struct {
	Server 	   *Server
	metadata   *trigger.Metadata
	runner     action.Runner
	config     *trigger.Config
	logger     log.Logger
	connection *eftl.Connection
	stop       chan bool
}

// New implements trigger.Factory.New
func (f *Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{metadata: f.Metadata(),config: config}, nil
}

// Init implements trigger.Init
func (t *Trigger) Initialize(ctx trigger.InitContext) error {
	addr := ":" + strconv.Itoa(DefaultPort)
	router := httprouter.New()

	for _, handler := range ctx.GetHandlers() {

		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}

		err = t.testActionHandler(handler)
		if err != nil {
			return err
		}
		//router.Handle("GET", "/a", t.newActionHandler(handler))
	}
	t.Server = NewServer(addr, router)
	return nil
}

func (t *Trigger) testActionHandler(handler trigger.Handler) error{
	fmt.Println("Inside Trigger action handler")
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	ca := t.config.Settings[settingCA]
	if ca != "" {
		certificate, err := ioutil.ReadFile(ca.(string))
		if err != nil {
			t.logger.Errorf("can't open certificate", err)
			return err
		}
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(certificate)
		tlsConfig = &tls.Config{
			RootCAs: pool,
		}
	}
	id := t.config.Settings[settingID]
	user := t.config.Settings[settingUser]
	password := t.config.Settings[settingPassword]
	fmt.Println("ID : ", id)
	options := &eftl.Options{
		ClientID:  id.(string),
		Username:  user.(string),
		Password:  password.(string),
		TLSConfig: tlsConfig,
	}

	url := t.config.Settings[settingURL]
	errorsChannel := make(chan error, 1)
	connectVal, err := eftl.Connect(url.(string), options, errorsChannel)
	if err != nil {
		t.logger.Errorf("connection failed: %s", err)
		return err
	}
	t.connection = connectVal

	messages := make(chan eftl.Message, 1000)
	dest := handler.Settings()
	matcher := fmt.Sprintf("{\"_dest\":\"%s\"}", dest[settingDest])
	_, err = t.connection.Subscribe(matcher, "", messages)
	if err != nil {
		t.logger.Errorf("subscription failed: %s", err)
		return err
	}
	t.stop = make(chan bool, 1)
	go func() {
		for {
			select {
			case message := <-messages:
				fmt.Println("Inside case")
				value := message["content"]
				content, ok := value.([]byte)
				fmt.Println("Content :", string(content))
				if !ok {
					content = []byte{}
				}
				var mimeMap map[string]interface{}
				err := util.Unmarshal("", message, &mimeMap)
				if err != nil {
					return err
				}
				fmt.Println("Mimemap:", mimeMap)
				/*out := t.constructStartRequest(w,r, ps,content)
				//results, err := handler.Handle(context.Background(), out)
				reply := &Reply{}
				reply.FromMap(results)

				if err != nil {
					t.logger.Debugf("Error: %s", err.Error())
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				if reply.Data != nil {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					if reply.Code == 0 {
						reply.Code = 200
					}
					w.WriteHeader(reply.Code)
					if err := json.NewEncoder(w).Encode(reply.Data); err != nil {
						log.Error(err)
					}
					return
				}

				if reply.Code > 0 {
					w.WriteHeader(reply.Code)
				} else {
					w.WriteHeader(http.StatusOK)
				}*/
			case err := <-errorsChannel:
				t.logger.Errorf("connection error: %s", err)
			case <-t.stop:
				fmt.Println("inside stop")
				return
			}
		}
	}()
	return nil
}


/*func (t *Trigger) newActionHandler(handler trigger.Handler) httprouter.Handle{
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("Inside Trigger action handler")
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		ca := t.config.Settings[settingCA]
		if ca != "" {
			certificate, err := ioutil.ReadFile(ca.(string))
			if err != nil {
				t.logger.Errorf("can't open certificate", err)
				return
			}
			pool := x509.NewCertPool()
			pool.AppendCertsFromPEM(certificate)
			tlsConfig = &tls.Config{
				RootCAs: pool,
			}
		}
		id := t.config.Settings[settingID]
		user := t.config.Settings[settingUser]
		password := t.config.Settings[settingPassword]
		fmt.Println("ID : ", id)
		options := &eftl.Options{
			ClientID:  id.(string),
			Username:  user.(string),
			Password:  password.(string),
			TLSConfig: tlsConfig,
		}

		url := t.config.Settings[settingURL]
		errorsChannel := make(chan error, 1)
		connectVal, err := eftl.Connect(url.(string), options, errorsChannel)
		if err != nil {
			t.logger.Errorf("connection failed: %s", err)
			return
		}
		t.connection = connectVal

		messages := make(chan eftl.Message, 1000)
		dest := handler.Settings()
		matcher := fmt.Sprintf("{\"_dest\":\"%s\"}", dest[settingDest])
		_, err = t.connection.Subscribe(matcher, "", messages)
		if err != nil {
			t.logger.Errorf("subscription failed: %s", err)
			return
		}
		t.stop = make(chan bool, 1)
		go func() {
			for {
				select {
				case _ = <-messages:
					fmt.Println("Inside case")

					out := t.constructStartRequest(w,r, ps)
					results, err := handler.Handle(context.Background(), out)
					reply := &Reply{}
					reply.FromMap(results)

					if err != nil {
						t.logger.Debugf("Error: %s", err.Error())
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}

					if reply.Data != nil {
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						if reply.Code == 0 {
							reply.Code = 200
						}
						w.WriteHeader(reply.Code)
						if err := json.NewEncoder(w).Encode(reply.Data); err != nil {
							log.Error(err)
						}
						return
					}

					if reply.Code > 0 {
						w.WriteHeader(reply.Code)
					} else {
						w.WriteHeader(http.StatusOK)
					}
				case err := <-errorsChannel:
					t.logger.Errorf("connection error: %s", err)
				case <-t.stop:
					fmt.Println("inside stop")
					return
				}
			}
		}()

	}
}*/


// Start implements ext.Trigger.Start
func (t *Trigger) Start() error {
	return t.Server.Start()
}

// Stop implements ext.Trigger.Stop
func (t *Trigger) Stop() error {
	if t.connection != nil {
		t.connection.Disconnect()
	}
	return t.Server.Stop()
}


func (t *Trigger) constructStartRequest(w http.ResponseWriter,r *http.Request, ps httprouter.Params) *Output {
	out := &Output{}

	out.PathParams = make(map[string]string)
	for _, param := range ps {
		out.PathParams[param.Key] = param.Value
	}

	queryValues := r.URL.Query()
	out.QueryParams = make(map[string]string, len(queryValues))

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
			t.logger.Errorf("Error while parsing query string: %s", err.Error())
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
	default:
		var content interface{}
		err := json.NewDecoder(r.Body).Decode(&content)
		if err != nil {
			switch {
			case err == io.EOF:
			// empty body
			//todo should handler say if content is expected?
			case err != nil:
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		}
		out.Content = content
	}
	return out
}
