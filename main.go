package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fmo/jobs/jobs"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var j []jobs.Job

	loader := jobs.NewLoader()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	content, err := loader.GetJobsFromFile(fmt.Sprintf("%s/%s", cwd, "jobs.json"))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &j)
	if err != nil {
		panic(err)
	}

	logChannel := make(chan jobs.Log)

	go func() {
		for log := range logChannel {
			fmt.Println(log.Message)
		}
	}()

	for _, job := range j {
		job.LogChannel = logChannel
		job.Start(ctx)
	}

	time.Sleep(35 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
