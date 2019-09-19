package tcpudp

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Network   string `md:"network"`       // The network type
	Host      string `md:"host"`          // The host name or IP for TCP server.
	Port      string `md:"port,required"` // The port to listen on
	Delimiter string `md:"delimiter"`     // Data delimiter for read and write
	TimeOut   int    `md:"timeout"`
}

type HandlerSettings struct {
}

type Output struct {
	Data string `md:"data"` // The data received from the connection
}

type Reply struct {
	Reply string `md:"reply"` // The reply to be sent back
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": o.Data,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Data, err = coerce.ToString(values["data"])
	if err != nil {
		return err
	}

	return nil
}

func (r *Reply) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"reply": r.Reply,
	}
}

func (r *Reply) FromMap(values map[string]interface{}) error {

	var err error
	r.Reply, err = coerce.ToString(values["reply"])
	if err != nil {
		return err
	}

	return nil
}
