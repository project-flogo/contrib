package main

import (
	"github.com/project-flogo/contrib/trigger/eftl/eftlHelpers"
)

func main() {
	errChannel := make(chan error, 1)
	options := &eftlHelpers.Options{
		ClientID: "test",
	}
	connection, err := eftlHelpers.Connect("ws://localhost:9191/channel", options, errChannel)
	if err != nil {
		panic(err)
	}
	defer connection.Disconnect()
	connection.Publish(eftlHelpers.Message{
		"_dest":   "sample",
		"content": []byte(`{"message": "hello world"}`),
	})
}
