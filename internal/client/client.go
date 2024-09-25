package client

import (
	"bufio"
	"fmt"
	"log"
	"os"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func InitClientConnection(username, password string) {

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["user"] = username
	uri := "http://localhost:8000/"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})

	client.On("connect", func() {
		fmt.Printf("connected\n")
		log.Printf("on connect\n")

	})

	client.On("message", func(msg string) {
		log.Printf("on message:%v\n", msg)

	})
	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)

		client.Emit("message", username+": "+command)
		log.Printf("You:%v\n", command)
	}
}
