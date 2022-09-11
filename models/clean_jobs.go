package models

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

func CleanJobs(max_age time.Duration) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("[error] failed to get current working directory '%s", err.Error())
	}
	var garbageJobs []string
	// filepath.WalkDir (go1.16) is more efficient thant filepath.Walk
	err = filepath.WalkDir(cwd, func(path string, d fs.DirEntry, err error) error {
		// Get all files in current directory
		if filepath.Ext(path) == ".pending" || filepath.Ext(path) == ".json" {
			garbageJobs = append(garbageJobs, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("[error] failed to walk dir %s", err.Error())
	}
	for _, jobFile := range garbageJobs {
		var job Job
		job, err = LoadJob(jobFile)
		if err != nil {
			log.Printf("[error] failed to read job from file: %s", err.Error())
			return
		}
		if err = deleteJob(job, max_age, jobFile); err != nil {
			log.Printf("[error] failed to update job %s: %s", job.Id, err.Error())
		}
	}
}
