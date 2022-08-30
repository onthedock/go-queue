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
	var max_age time.Duration
	var loop bool
	flag.DurationVar(&interval, "interval", 3*time.Second, "When running with -loop, interval between every check for outdated jobs")
	flag.DurationVar(&max_age, "max_age", 1*time.Hour, "Remove jobs older than max_age")
	flag.BoolVar(&loop, "loop", false, "Run the cleaner app in a loop")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	if loop {
		// Main loop
		log.Printf("cleaning jobs older than '%s' every '%s' ...\n", max_age.String(), interval.String())
		for {
			select {
			case <-c:
				log.Println("Exiting the app...")
				return
			case <-time.After(interval):
				log.Printf("cleaning jobs older than '%s' every '%s' ...\n", max_age.String(), interval.String())
				models.CleanJobs(max_age)
			}
		}
	}
	log.Printf("[info] deleting jobs older than '%s'\n", max_age.String())
	models.CleanJobs(max_age)
}
