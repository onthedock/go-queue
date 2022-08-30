package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"queue/models"
	"syscall"
	"time"
)

func main() {
	var interval time.Duration
	flag.DurationVar(&interval, "interval", 2*time.Second, "Interval between every check for pending jobs")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Main loop
	log.Printf("processing pending jobs every %s ...\n", interval.String())
	for {
		select {
		case <-c:
			log.Println("Exiting the app...")
			return
		case <-time.After(interval):
			log.Printf("processing pending jobs every %s ...\n", interval.String())
			models.PendingJobs()
		}
	}
}
