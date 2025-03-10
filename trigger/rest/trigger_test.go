package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/api"
	"github.com/project-flogo/core/engine"
	"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/support/test"
	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

func TestTrigger_Register(t *testing.T) {

	ref := support.GetRef(&Trigger{})
	f := trigger.GetFactory(ref)
	assert.NotNil(t, f)
}

const testConfig string = `{
	"id": "trigger-restOld",
	"ref": "github.com/project-flogo/contrib/trigger/restOld",
	"settings": {
        "port": "8888"
    },
	"handlers": [
	  {
		"settings": {
		  "isPassThroughUri": "NO",
		  "method": "GET",
		  "path": "/test"
		},
		"action" : {
		  "id": "dummy"
		}
	  }
	]
  }`

func TestRestTrigger_Initialize(t *testing.T) {

	f := &Factory{}

	config := &trigger.Config{}
	err := json.Unmarshal([]byte(testConfig), config)
	assert.Nil(t, err)

	actions := map[string]action.Action{"dummy": test.NewDummyAction(func() {
	})}

	trg, err := test.InitTrigger(f, config, actions)

	assert.Nil(t, err)
	assert.NotNil(t, trg)
	err = trg.Start()
	assert.Nil(t, err)
	err = trg.Stop()
	assert.Nil(t, err)

}

func Test_App(t *testing.T) {
	var wg sync.WaitGroup
	app := myApp("NO")

	e, err := api.NewEngine(app)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//assert.Nil(t, err)

	wg.Add(1)
	go engine.RunEngine(e)

	go func() {
		time.Sleep(5 * time.Second)
		bodyReader := strings.NewReader("{\"marker\":\"hello\"}")
		_, err := http.Post("http://localhost:5050", "application/json", bodyReader)
		if err != nil {
			panic("failed to connect: " + err.Error())
		}

		if err != nil {
			assert.NotNil(t, err)
			wg.Done()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("The response is")
	e.Stop()
}

func Test_AppWith(t *testing.T) {
	var wg sync.WaitGroup
	app := myApp("YES")

	e, err := api.NewEngine(app)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//assert.Nil(t, err)

	wg.Add(1)
	go engine.RunEngine(e)

	go func() {
		time.Sleep(5 * time.Second)
		bodyReader := strings.NewReader("{\"marker\":\"hello\"}")
		_, err := http.Post("http://localhost:5050", "application/json", bodyReader)
		if err != nil {
			panic("failed to connect: " + err.Error())
		}

		if err != nil {
			assert.NotNil(t, err)
			wg.Done()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("The response is")
	e.Stop()
}

func myApp(isProxyValue string) *api.App {

	app := api.NewApp()

	trg := app.NewTrigger(&Trigger{}, &Settings{Port: 5050})

	h, _ := trg.NewHandler(&HandlerSettings{IsPassThroughUri: isProxyValue, Method: "GET", Path: "/test"})

	h.NewAction(RunActivities)

	return app

}

func RunActivities(ctx context.Context, inputs map[string]interface{}) (map[string]interface{}, error) {

	result := &Reply{Code: 200, Data: "hello"}
	return result.ToMap(), nil
}
