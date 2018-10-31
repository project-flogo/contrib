package ping

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
	"github.com/project-flogo/contrib/tree/master/trigger/ping"
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
}

// Factory Ping Trigger factory
type Factory struct {}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return pingTriggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	response := metadata.PingResponse{
		Version:        config.GetSetting("version"),
		Appversion:     config.GetSetting("appversion"),
		Appdescription: config.GetSetting("appdescription"),
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Error("Ping service data formation error")
	}

	port := config.GetSetting("port")
	if len(port) == 0 {
		port = DefaultPort
	}

	mux := http.NewServeMux()
	trigger := &Trigger{
		metadata: f.metadata,
		config:   config,
		response: string(data),
		Server: &http.Server{
			Addr:    ":" + port,
			Handler: mux,
		},
	}

	mux.HandleFunc("/ping", trigger.PingResponseHandlerShort)
	mux.HandleFunc("/ping/details", trigger.PingResponseHandlerDetail)
	return trigger,nil
}


// Initialize start the Ping service
func (t *Trigger) Initialize(context trigger.InitContext) error {
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
	log.Info("Ping service starting...")

	go func() {
		if err := t.ListenAndServe(); err != http.ErrServerClosed {
			log.Errorf("Ping service err:", err)
		}
	}()
	log.Info("Ping service started")
	return nil
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	if err := t.Shutdown(nil); err != nil {
		log.Errorf("[mashling-ping-service] Ping service error when stopping:", err)
		return err
	}
	log.Info("[mashling-ping-service] Ping service stopped")
	return nil
}
