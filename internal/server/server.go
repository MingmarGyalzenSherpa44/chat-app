package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/MingmarGyalzenSherpa44/chat-app/internal/database"
	socketio "github.com/googollee/go-socket.io"
)

func InitServer() {

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")

		s.Join("our-room")
		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {

		//parse query
		query := s.URL().RawQuery
		parsedQuery, err := url.ParseQuery(query)

		if err != nil {
			log.Fatal("Error parsing query")
		}

		userId := parsedQuery.Get("user_id")
		parsedUserId, err := strconv.Atoi(userId)
		if err != nil {
			log.Fatal("error parsing user id")
		}
		database.SaveMessage(parsedUserId, msg)

		user := parsedQuery.Get("user")
		fmt.Println(parsedQuery)

		server.BroadcastToRoom("/", "our-room", "message", fmt.Sprintf("%v : %v", user, msg))
		fmt.Printf("Broadcasted")

	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {

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
