package models

import (
	"encoding/json"
	"log"
	"os"
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
