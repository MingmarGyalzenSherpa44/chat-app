package server

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func InitServer() {

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())

		s.Join("our-room")

		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {

		fmt.Println(msg)

		server.BroadcastToRoom("/", "our-room", "message", msg)
		fmt.Printf("Broadcasted")

	})

	server.OnError("/", func(s socketio.Conn, e error) {
		// server.Remove(s.ID())
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		// Add the Remove session id. Fixed the connection & mem leak
		// server.Remove(s.ID())
		s.Leave("our-room")
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("///////////////////////////////")
	log.Println("////////SERVER STARTED///////")
	log.Println("///////////////////////////////")
	log.Println("Listening at :8000")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
