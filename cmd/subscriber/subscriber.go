// https://oswalt.dev/2019/09/kicking-the-tires-with-the-nats-go-client/

package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	log.Printf("Connecting to nats.DefautlURL: %s", nats.DefaultURL)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("unable to connect to defaultURL: %s", err.Error())
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("unable to wrap connection with JSON encoder %s", err.Error())
	}

	defer ec.Close()

	log.Println("Connected to NATS and ready to receive messages")
	type Request struct {
		Id int
	}

	personChanRecv := make(chan *Request)
	ec.BindRecvChan("topic", personChanRecv)

	for {
		req := <-personChanRecv

		log.Printf("received request %d", req.Id)
	}
}
