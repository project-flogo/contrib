package ping

import (
	"fmt"
	"strconv"
	"encoding/json"
	"io"
	"net/http"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
	"runtime"
	"net/trace"
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

	port := strconv.Itoa(config.Settings["port"].(int))
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
	fmt.Println("After init :")
	PrintMemUsage()
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
	tr := trace.New("TraceTest", req.URL.Path)
	defer tr.Finish()
	PrintMemUsage()
	io.WriteString(w, t.response+"\n")
	tr.LazyPrintf("Details through trace")
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
	var rtm runtime.MemStats
	var m Monitor
	runtime.ReadMemStats(&rtm)

	// Number of goroutines
	m.NumGoroutine = runtime.NumGoroutine()

	// Misc memory stats
	m.Alloc = rtm.Alloc
	m.TotalAlloc = rtm.TotalAlloc
	m.Sys = rtm.Sys
	m.Mallocs = rtm.Mallocs
	m.Frees = rtm.Frees

	// Live objects = Mallocs - Frees
	m.LiveObjects = m.Mallocs - m.Frees

	// GC Stats
	m.PauseTotalNs = rtm.PauseTotalNs
	m.NumGC = rtm.NumGC

	// Just encode to json and print
	b, _ := json.Marshal(m)
	fmt.Println(string(b))
}


func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

type Monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs uint64

	NumGC        uint32
	NumGoroutine int
}