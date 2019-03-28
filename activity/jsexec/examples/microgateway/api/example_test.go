package api

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/project-flogo/core/engine"
	"github.com/project-flogo/microgateway/api"
	"github.com/stretchr/testify/assert"
)

func Drain(port string) {
	for {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort("", port), time.Second)
		if conn != nil {
			conn.Close()
		}
		if err != nil && strings.Contains(err.Error(), "connect: connection refused") {
			break
		}
	}
}

func Pour(port string) {
	for {
		conn, _ := net.Dial("tcp", net.JoinHostPort("", port))
		if conn != nil {
			conn.Close()
			break
		}
	}
}

// Response is a reply form the server
type Response struct {
	Sum float64 `json:"sum"`
}

func testApplication(t *testing.T, e engine.Engine) {
	defer api.ClearResources()
	Drain("9096")
	err := e.Start()
	assert.Nil(t, err)
	defer func() {
		err := e.Stop()
		assert.Nil(t, err)
	}()
	Pour("9096")

	transport := &http.Transport{
		MaxIdleConns: 1,
	}
	defer transport.CloseIdleConnections()
	client := &http.Client{
		Transport: transport,
	}
	request := func() Response {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:9096/calculate", nil)
		assert.Nil(t, err)
		response, err := client.Do(req)
		assert.Nil(t, err)
		body, err := ioutil.ReadAll(response.Body)
		assert.Nil(t, err)
		response.Body.Close()
		rsp := Response{}
		err = json.Unmarshal(body, &rsp)
		assert.Nil(t, err)
		return rsp
	}

	response := request()
	assert.Equal(t, 3.0, response.Sum)
}

func TestIntegrationAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping API integration test in short mode")
	}

	e, err := Example()
	assert.Nil(t, err)
	testApplication(t, e)
}

func TestIntegrationJSON(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping JSON integration test in short mode")
	}

	data, err := ioutil.ReadFile(filepath.FromSlash("../json/flogo.json"))
	assert.Nil(t, err)
	cfg, err := engine.LoadAppConfig(string(data), false)
	assert.Nil(t, err)
	e, err := engine.New(cfg)
	assert.Nil(t, err)
	testApplication(t, e)
}
