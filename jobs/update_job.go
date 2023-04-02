package jobs

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func update(job Job) error {
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
