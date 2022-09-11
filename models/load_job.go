package models

import (
	"encoding/json"
	"log"
	"os"
)

func LoadJob(jobFile string) (Job, error) {
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
