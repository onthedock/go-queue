package main

import (
	"encoding/json"
	"log"
	"os"
	"queue/models"
	"time"
)

func main() {
	var filename string = "89e47a15-50ec-4595-928e-d82503ee616c.json.pending"
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var j = new(models.Job)
	if err := json.Unmarshal(fileBytes, j); err != nil {
		log.Fatalf("[ERROR] Failed parse file '%s': '%s'", filename, err.Error())
	}
	updateJob(*j)
}

func updateJob(job models.Job) error {
	job.LastUpdated = time.Now()
	job.Result = job.Num1 + job.Num2
	j, err := json.Marshal(job)
	if err != nil {
		log.Fatalf("[ ERROR ] Failed to convert struct  to JSON '%s'\n", err.Error())
		return err
	}
	if err := os.WriteFile(job.Id+".json", []byte(j), 0664); err != nil {
		log.Fatalf("[ERROR] Failed to update job '%s', '%s'", job.Id, err.Error())
		return err
	}
	if err := os.Remove(job.Id + ".json.pending"); err != nil {
		log.Printf("[ERROR] Failed to delete file '%s', '%s'", job.Id+".json.pending", err.Error())
		return err
	}
	return nil
}
