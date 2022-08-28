package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Main loop
	for {
		select {
		case <-c:
			log.Println("Exiting the app...")
			return
		case <-time.After(1 * time.Second):
			log.Println("processing pending jobs...")
			pendingJobs()
		}
	}

	// var filename string = "907292c5-b00e-4133-ae81-492743177605.json.pending"
	// fileBytes, err := os.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// var j = new(models.Job)
	// if err := json.Unmarshal(fileBytes, j); err != nil {
	// 	log.Fatalf("[ERROR] Failed parse file '%s': '%s'", filename, err.Error())
	// }
	// updateJob(*j)
}

func pendingJobs() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("[error] Failed to get current working directory '%s", err.Error())
	}
	// var pendingJobs []string
	// filepath.WalkDir (go1.16) is more efficient thant filepath.Walk
	err = filepath.WalkDir(cwd, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".pending" {
			fmt.Printf("path %s\n", path)
		}

		return nil
	})
	if err != nil {
		log.Fatalf("[error] Failed walk dir %s", err.Error())
	}

}

// func updateJob(job models.Job) error {
// 	job.LastUpdated = time.Now()
// 	job.Result = job.Num1 + job.Num2
// 	j, err := json.Marshal(job)
// 	if err != nil {
// 		log.Fatalf("[ ERROR ] Failed to convert struct  to JSON '%s'\n", err.Error())
// 		return err
// 	}
// 	if err := os.WriteFile(job.Id+".json", []byte(j), 0664); err != nil {
// 		log.Fatalf("[ERROR] Failed to update job '%s', '%s'", job.Id, err.Error())
// 		return err
// 	}
// 	if err := os.Remove(job.Id + ".json.pending"); err != nil {
// 		log.Printf("[ERROR] Failed to delete file '%s', '%s'", job.Id+".json.pending", err.Error())
// 		return err
// 	}
// 	return nil
// }
