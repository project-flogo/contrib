package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	errChannel := make(chan error, 1)
	options := &Options{
		ClientID: "test",
	}
	connection, err := Connect("ws://localhost:9191/channel", options, errChannel)
	if err != nil {
		panic(err)
	}
	defer connection.Disconnect()
	connection.Publish(Message{
		"_dest":   "sample",
		"content": []byte(`{"message": "hello world"}`),
	})
}
