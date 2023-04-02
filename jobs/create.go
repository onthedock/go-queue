package jobs

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func Create(n1, n2 int) (uuid.UUID, error) {
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
