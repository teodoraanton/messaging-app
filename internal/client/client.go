package client

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"wantsome.ro/messagingapp/pkg/models"
)

func RunClient() {
	url := "ws://localhost:8080/ws"
	randId := rand.Intn(10)
	message := models.Message{Message: fmt.Sprintf("Hello world from my client %d !", randId), UserName: fmt.Sprintf("Client %d", randId)}

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatalf("error dialing %s\n", err)
	}
	defer c.Close()

	done := make(chan bool)
	// reading server messages
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("error reading: %s\n", err)
				return
			}
			fmt.Printf("Got message: %s\n", message)
		}
	}()

	// writing messages to server
	go func() {
		for {
			err := c.WriteJSON(message)
			if err != nil {
				log.Printf("error writing %s\n", err)
				return
			}
			time.Sleep(3 * time.Second)
		}
	}()

	<-done
}
