package ping

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

// DefaultPort is the default port for Ping service
const DefaultPort = "9096"

type Settings struct {
	Port 		int 	`md:"port,required"`
	Version 	string 	`md:"version"`
	AppVersion 	string 	`md:"appversion"`
	AppDescription 	string 	`md:"appdescription"`
}

var triggerMd = trigger.NewMetadata(&Settings{})

func init() {
	trigger.Register(&Trigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// Trigger is the ping trigger
type Trigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	response string
	*http.Server
	logger   log.Logger
}

// New implements trigger.Factory.New
func (f *Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	type PingResponse struct {
		Version        string
		Appversion     string
		Appdescription string
	}
	fmt.Println("config:", config)
	response := PingResponse{
		Version:        "version",//config.Settings.Version,
		Appversion:     "appversion",//config.Settings.AppVersion,
		Appdescription: "appdescr",//config.Settings.AppDescription,
	}

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Ping service data formation error")
	}

	port := DefaultPort//config.Settings.Port
	if len(port) == 0 {
		port = DefaultPort
	}

	mux := http.NewServeMux()
	trigger := &Trigger{
		metadata: f.Metadata(),
		config:   config,
		response: string(data),
		Server: &http.Server{
			Addr:    ":" + port,
			Handler: mux,
		},
	}
	fmt.Println("trigger metadata is :", trigger.metadata)
	fmt.Println("trigger config is :", trigger.config)
	fmt.Println("trigger response is :", trigger.response)
	fmt.Println("trigger server is :", trigger.Server)
	mux.HandleFunc("/ping", trigger.PingResponseHandlerShort)
	mux.HandleFunc("/ping/details", trigger.PingResponseHandlerDetail)
	return trigger, nil
}

// Init implements trigger.Init
func (t *Trigger) Initialize(ctx trigger.InitContext) error {
	t.Start()
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
