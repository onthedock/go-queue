package jobs

import "time"

type Job struct {
	Id          string    `json:"id"`
	Num1        int       `json:"num1"`
	Num2        int       `json:"num2"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"updated"`
	Result      int       `json:"result"`
}
