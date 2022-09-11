package models

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func PendingJobs() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("[error] Failed to get current working directory '%s", err.Error())
	}
	var pendingJobs []string
	// filepath.WalkDir (go1.16) is more efficient thant filepath.Walk
	err = filepath.WalkDir(cwd, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".pending" {
			pendingJobs = append(pendingJobs, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("[error] Failed walk dir %s", err.Error())
	}
	for _, j := range pendingJobs {
		var job Job
		job, err = LoadJob(j)
		if err != nil {
			log.Printf("[error] failed to read job from file: %s", err.Error())
			return
		}
		if err = updateJob(job); err != nil {
			log.Printf("[error] failed to update job %s: %s", job.Id, err.Error())
		}
	}
}
