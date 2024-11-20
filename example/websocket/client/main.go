package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	serverAddr := "ws://localhost:8080/ws"

	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("fail to connect: ", err)
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("ReadMessage error: ", err)
				return
			}
			log.Printf("received: %s\n", message)
		}
	}()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte("{\"username\": \"user1\", \"message\": \"Hello\"}"))
			if err != nil {
				log.Println("WriteMessage error: ", err)
				return
			}
		case <-interrupt:
			log.Println("Interrupt received, closing connection")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Close error:", err)
				return
			}
			return
		}
	}
}
