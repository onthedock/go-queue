// https://oswalt.dev/2019/09/kicking-the-tires-with-the-nats-go-client/

package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("fatal error connecting to NATS server %s", err.Error())
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("unable to use the JSON encoder for the connection %s", err.Error())
	}

	defer ec.Close()

	log.Println("Connected to NATS and ready to send messages")

	type Request struct {
		Id int
	}

	personChanSend := make(chan *Request)
	ec.BindSendChan("topic", personChanSend)

	i := 0
	for {
		// Create instance of type Request with Id set to current value of i
		req := Request{Id: i}

		log.Printf("Sending request %d", req.Id)
		personChanSend <- &req

		// Pause and increment the counter
		time.Sleep(1 * time.Second)
		i = i + 1
	}
}
