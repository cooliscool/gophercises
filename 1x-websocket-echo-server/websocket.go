// websockets.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
}

func main() {

	const bigMessage string = "small msg"

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// log the incoming request
		log.Println("Incoming Request")
		log.Println("Header: ")
		log.Println(r.Header)
		log.Println("Body: ")
		log.Println(r.Body)

		headers := make(http.Header)

		if ua := r.Header.Get("User-Agent"); ua != "" {
			headers.Add("Echo-User-Agent", r.Header.Get("User-Agent"))
		}

		conn, _ := upgrader.Upgrade(w, r, headers) // error ignored for sake of simplicity

		normalMessage := 1

		if err := conn.WriteMessage(normalMessage, []byte(bigMessage)); err != nil {
			return
		}

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ws.html")
	})

	fmt.Println("Starting the server on : ", 8888)
	fmt.Println("Hit localhost:8888 in browser for communicating with echo server running at localhost:8888/echo")
	http.ListenAndServe(":8888", nil)
}
