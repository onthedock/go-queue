package models

import "time"

type Job struct {
	Id          int       `json:"id"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"updated"`
	Status      string    `json:"status"`
	Result      string    `json:"result"`
}

var jobs []Job = make([]Job, 0)

func initialize() {
	for i := 0; i < 100; i++ {
		job := Job{i, time.Now(), time.Now(), "pending", "none"}
		jobs = append(jobs, job)
	}

}

func AddJob(n1, n2 int) (int, error) {
	return n1 + n2, nil
}

func GetJob(jid int) (Job, error) {
	initialize()
	return jobs[jid], nil
}
