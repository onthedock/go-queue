package main

import (
	"log"
	"os"
	"os/signal"
	"queue/models"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Main loop
	for {
		select {
		case <-c:
			log.Println("Exiting the app...")
			return
		case <-time.After(5 * time.Second):
			log.Println("processing pending jobs...")
			models.PendingJobs()
		}
	}
}
