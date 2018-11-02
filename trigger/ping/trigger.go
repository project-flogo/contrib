package ping

import (
	"fmt"
	"strconv"
	"encoding/json"
	"io"
	"net/http"
	"github.com/project-flogo/core/trigger"
)

// DefaultPort is the default port for Ping service
const DefaultPort = "9096"

var triggerMd = trigger.NewMetadata(&Settings{})
var statsResult string

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
}

// New implements trigger.Factory.New
func (f *Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	type PingResponse struct {
		Version        	string
		Appversion     	string
		Appdescription 	string
	}
	response := PingResponse{
		Version:        config.Settings["version"].(string),
		Appversion:     config.Settings["appversion"].(string),
		Appdescription: config.Settings["appdescription"].(string),
	}

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Errorf("Ping service data formation error",err)
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
	statsResult = PrintMemUsage()
	return trigger, nil
}

// Init implements trigger.Init
func (t *Trigger) Initialize(ctx trigger.InitContext) error {
	return nil
}

func (t *Trigger) PingResponseHandlerShort(w http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("Authorization")
	if(Valid(token)) {
		io.WriteString(w, "{\"response\":\"Ping successful\"}\n")
	}else{
		io.WriteString(w, "Error : Invalid User !! \n")
	}
}

//PingResponseHandlerDetail handles simple response
func (t *Trigger) PingResponseHandlerDetail(w http.ResponseWriter, req *http.Request) {
	token := t.config.Settings["password"].(string)
	if(Valid(token)) {
		io.WriteString(w, t.response + "\n")
		io.WriteString(w, "Details :\n")
		io.WriteString(w, statsResult + "\n")
	}else{
		io.WriteString(w, "Error : Invalid User !! \n")
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

// Start implements util.Managed.Start
func (t *Trigger) Start() error {
	go func() {
		if err := t.Server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Errorf("Ping service err:", err)
		}
	}()
	return nil
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	if err := t.Server.Shutdown(nil); err != nil {
		fmt.Errorf("[mashling-ping-service] Ping service error when stopping:", err)
		return err
	}
	return nil
}

