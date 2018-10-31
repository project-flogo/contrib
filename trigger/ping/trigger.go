package ping

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

const (
	// DefaultPort is the default port for Ping service
	DefaultPort = "9090"
)

var pingTriggerMd = trigger.NewMetadata(&Settings{})

func init() {
	trigger.Register(&Trigger{}, &Factory{})
}

// Trigger is the ping trigger
type Trigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	response string
	*http.Server
	logger   log.Logger
}

// Factory Ping Trigger factory
type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (f *Factory) Metadata() *trigger.Metadata {
	return pingTriggerMd
}

// New implements trigger.Factory.New
func (f *Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}
	return &Trigger{metadata: f.Metadata(), config:config}, nil
}


// Initialize start the Ping service
func (t *Trigger) Initialize(context trigger.InitContext) error {
	t.logger = context.Logger()
	response := PingResponse{
		Version:        t.config.GetSetting("version"),
		Appversion:     t.config.GetSetting("appversion"),
		Appdescription: t.config.GetSetting("appdescription"),
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.logger.Errorf("Ping service data formation error")
	}

	port := t.config.GetSetting("port")
	if len(port) == 0 {
		port = DefaultPort
	}

	mux := http.NewServeMux()
	t.response = string(data)
	t.Server = &http.Server{Addr:    ":" + port, Handler: mux, }

	mux.HandleFunc("/ping", trigger.PingResponseHandlerShort)
	mux.HandleFunc("/ping/details", trigger.PingResponseHandlerDetail)
	return nil
}

func (t *Trigger) PingResponseHandlerShort(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "{\"response\":\"Ping successful\"}\n")
}

//PingResponseHandlerDetail handles simple response
func (t *Trigger) PingResponseHandlerDetail(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, t.response+"\n")
}

// Start implements util.Managed.Start
func (t *Trigger) Start() error {
	t.logger.Infof("Ping service starting...")

	go func() {
		if err := t.ListenAndServe(); err != http.ErrServerClosed {
			t.logger.Errorf("Ping service err:", err)
		}
	}()
	t.logger.Infof("Ping service started")
	return nil
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	if err := t.Shutdown(nil); err != nil {
		t.logger.Errorf("[mashling-ping-service] Ping service error when stopping:", err)
		return err
	}
	t.logger.Infof("[mashling-ping-service] Ping service stopped")
	return nil
}
