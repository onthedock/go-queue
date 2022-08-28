package models

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Job struct {
	Id          string    `json:"id"`
	Num1        int       `json:"num1"`
	Num2        int       `json:"num2"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"updated"`
	Result      int       `json:"result"`
}

func CreateJob(n1, n2 int) (uuid.UUID, error) {
	jobid := uuid.New()
	var job = Job{
		Id:          jobid.String(),
		Num1:        n1,
		Num2:        n2,
		Created:     time.Now(),
		LastUpdated: time.Now(),
		Result:      0,
	}
	j, err := json.Marshal(job)
	if err != nil {
		log.Fatalf("[ ERROR ] Failed to convert struct  to JSON '%s'\n", err.Error())
	}
	os.WriteFile(jobid.String()+".pending", []byte(j), 0664)

	return jobid, nil
}

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
		job, err = loadJob(j)
		if err != nil {
			log.Printf("[error] failed to read job from file: %s", err.Error())
			return
		}
		if err = updateJob(job); err != nil {
			log.Printf("[error] failed to update job %s: %s", job.Id, err.Error())
		}
	}
}

func loadJob(jobFile string) (Job, error) {
	var job = Job{}

	jobBytes, err := os.ReadFile(jobFile)
	if err != nil {
		log.Printf("[error] failed to read file %s: %s", jobFile, err.Error())
		return job, err
	}
	if err := json.Unmarshal(jobBytes, &job); err != nil {
		log.Printf("[error] failed to parse file '%s': '%s'", jobFile, err.Error())
		return job, err
	}
	return job, nil
}

func updateJob(job Job) error {
	job.LastUpdated = time.Now()
	job.Result = job.Num1 + job.Num2
	j, err := json.Marshal(job)
	if err != nil {
		log.Fatalf("[ ERROR ] Failed to convert struct to JSON '%s'\n", err.Error())
		return err
	}
	if err := os.WriteFile(job.Id+".json", []byte(j), 0664); err != nil {
		log.Fatalf("[ERROR] Failed to update job '%s', '%s'", job.Id, err.Error())
		return err
	}
	if err := os.Remove(job.Id + ".pending"); err != nil {
		log.Printf("[ERROR] Failed to delete file '%s', '%s'", job.Id+".pending", err.Error())
		return err
	}
	log.Printf("[ ok ] processed jobId '%s' successfully\n", job.Id)
	return nil
}
