package ping

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
	"runtime"
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
	Server *http.Server
	logger   log.Logger
}

// New implements trigger.Factory.New
func (f *Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	type PingResponse struct {
		Version        string
		Appversion     string
		Appdescription string
	}
	fmt.Println("config:", config.Settings)
	response := PingResponse{
		Version:        config.Settings["version"].(string),
		Appversion:     "",//config.Settings.AppVersion,
		Appdescription: "",//config.Settings.AppDescription,
	}

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Ping service data formation error")
	}

	port := config.Settings["port"].(string)
	if len(port) == 0 {
		port = DefaultPort
	}

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	trigger := &Trigger{
		metadata: f.Metadata(),
		config:   config,
		response: string(data),
		Server: server,
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
	return nil
}

func (t *Trigger) PingResponseHandlerShort(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "{\"response\":\"Ping successful\"}\n")
}

//PingResponseHandlerDetail handles simple response
func (t *Trigger) PingResponseHandlerDetail(w http.ResponseWriter, req *http.Request) {
	PrintMemUsage()
	io.WriteString(w, t.response+"\n")
}

// Start implements util.Managed.Start
func (t *Trigger) Start() error {
	fmt.Println("Inside trigger start")
	fmt.Println("Server:", t.Server)

	fmt.Println("Ping service starting...")

	go func() {
		fmt.Println("inside go routine")
		if err := t.Server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Errorf("Ping service err:", err)
		}
	}()
	fmt.Println("Ping service started")
	return nil
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	if err := t.Server.Shutdown(nil); err != nil {
		fmt.Errorf("[mashling-ping-service] Ping service error when stopping:", err)
		return err
	}
	fmt.Println("[mashling-ping-service] Ping service stopped")
	return nil
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Printf("No of GoRoutines active", runtime.NumGoroutine())
}


func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
