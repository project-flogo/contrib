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
		Version        	string
		Appversion     	string
		Appdescription 	string
		Endpoint	string
		NumGoroutine 	int
		Alloc,
		TotalAlloc,
		Sys,
		Mallocs,
		Frees,
		LiveObjects	uint64
		NumGC		uint32
	}
	fmt.Println("config:", config.Settings)
	response := PingResponse{
		Version:        config.Settings["version"].(string),
		Appversion:     config.Settings["appversion"].(string),
		Appdescription: config.Settings["appdescription"].(string),
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
	mux.HandleFunc("/ping", trigger.PingResponseHandlerShort)
	mux.HandleFunc("/ping/details", trigger.PingResponseHandlerDetail)
	trigger.PrintMemUsage()
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
	token := req.Header.Get("Authorization")
	fmt.Println(token)
	if(valid(token)) {
		t.response.Endpoint = req.URL.Path
		io.WriteString(w, t.response + "\n")
	}else{
		fmt.Errorf("Invalid token!!!")
	}

	//Another way to get trace : more details
	/*
	tr := trace.New("TraceTest", req.URL.Path)
	defer tr.Finish()
	fmt.Println(reflect.TypeOf(tr).String())
	fmt.Println("Trace:")
	fmt.Printf("%+v\n", tr)
	fmt.Println(tr)*/
}

func valid(token string) bool{
	matched, _ := regexp.MatchString("^Bearer (.*)", token)
	return matched
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

func (t *Trigger) PrintMemUsage() {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)

	// Number of goroutines
	t.NumGoroutine = runtime.NumGoroutine()

	// Misc memory stats
	t.Alloc = rtm.Alloc
	t.TotalAlloc = rtm.TotalAlloc
	t.Sys = rtm.Sys
	t.Mallocs = rtm.Mallocs
	t.Frees = rtm.Frees

	// Live objects = Mallocs - Frees
	t.LiveObjects = t.Mallocs - t.Frees

	//GC stats
	t.NumGC = rtm.NumGC

}
