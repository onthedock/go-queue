package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"queue/jobs"
	"syscall"
	"time"
)

func main() {
	var interval time.Duration
	flag.DurationVar(&interval, "interval", 2*time.Second, "Interval between every check for pending jobs")
	flag.Parse()

	catchCtrlC := make(chan os.Signal, 1)
	signal.Notify(catchCtrlC, os.Interrupt, syscall.SIGTERM)

	// Main loop
	log.Printf("processing pending jobs every %s ...\n", interval.String())
	for {
		select {
		case <-catchCtrlC:
			log.Println("Exiting the app...")
			return
		case <-time.After(interval):
			log.Printf("processing pending jobs every %s ...\n", interval.String())
			jobs.Pending()
		}
	}
}
