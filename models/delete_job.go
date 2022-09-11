package models

import (
	"log"
	"os"
	"time"
)

func deleteJob(job Job, max_age time.Duration, jobFile string) error {
	var age time.Duration = time.Since(job.LastUpdated)
	if age > max_age {
		if err := os.Remove(jobFile); err != nil {
			log.Printf("[error] failed to delete file '%s', %s", jobFile, err.Error())
			return err
		}
		log.Printf("[ ok ] deleted job '%s' (last modified '%s' ago)\n", job.Id, age.String())
	}
	return nil
}
