package rest

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/project-flogo/core/api"
	"github.com/project-flogo/core/engine"
	"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/trigger"
	"github.com/stretchr/testify/assert"
)

func TestTrigger_Register(t *testing.T) {

	ref := support.GetRef(&Trigger{})
	f := trigger.GetFactory(ref)
	assert.NotNil(t, f)
}
func Test_App(t *testing.T) {
	var wg sync.WaitGroup
	app := myApp()

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
		roots := x509.NewCertPool()

		conn, err := tls.Dial("tcp", "localhost:5050", &tls.Config{
			RootCAs: roots,
		})
		if err != nil {
			panic("failed to connect: " + err.Error())
		}
		conn.Close()
		if err != nil {
			assert.NotNil(t, err)
			wg.Done()
		}
		//assert.Equal(t, "text/plain; charset=UTF-8", resp.Header.Get("Content-type"))
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("The response is")
}
func myApp() *api.App {

	app := api.NewApp()

	trg := app.NewTrigger(&Trigger{}, &Settings{Port: 5050, TLS: true, CertPm: "/Users/skothari-tibco/Desktop/cert.pem", KeyPm: "/Users/skothari-tibco/Desktop/key.pem"})

	h, _ := trg.NewHandler(&HandlerSettings{Method: "GET", Path: "/test"})

	h.NewAction(RunActivities)

	return app

}
func RunActivities(ctx context.Context, inputs map[string]interface{}) (map[string]interface{}, error) {

	result := &Reply{Code: 200, Data: "hello"}
	return result.ToMap(), nil
}
