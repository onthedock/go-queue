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
	var loop time.Duration
	var max_age time.Duration
	flag.DurationVar(&loop, "interval", 0, "If a value is provided, run every (value); if the value is 0, run just once")
	flag.DurationVar(&max_age, "max_age", 1*time.Hour, "Remove jobs older than max_age")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	if loop != 0 {
		// Main loop
		log.Printf("cleaning jobs older than '%s' every '%s' ...\n", max_age.String(), loop.String())
		for {
			select {
			case <-c:
				log.Println("Exiting the app...")
				return
			case <-time.After(loop):
				log.Printf("cleaning jobs older than '%s' every '%s' ...\n", max_age.String(), loop.String())
				models.CleanJobs(max_age)
			}
		}
	}
	log.Printf("[info] deleting jobs older than '%s'\n", max_age.String())
	models.CleanJobs(max_age)
}
