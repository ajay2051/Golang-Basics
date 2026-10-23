package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Value string
	Err   error
}

func worker(jobsChannel chan string, wg *sync.WaitGroup, resultChannel chan Result) string {

	defer wg.Done()

	for job := range jobsChannel {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("Image Processed: %s\n", job)
		resultChannel <- Result{
			Value: job,
			Err:   nil,
		}
	}
	fmt.Println("Worker Shutting down...")
	return <-jobsChannel

}

func main() {

	jobs := []string{
		"image_1.png",
		"image_2.png",
		"image_3.png",
		"image_4.png",
		"image_5.png",
	}

	var wg sync.WaitGroup
	totalWorkers := 5

	resultChannel := make(chan Result, 2)
	jobsChannel := make(chan string, len(jobs))

	startTime := time.Now()

	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(jobsChannel, &wg, resultChannel)
	}

	// wg.Add(2)
	//
	// go worker("image_1.png", &wg, resultChannel)
	// go worker("image_2.png", &wg, resultChannel)
	//

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// send the jobs

	for _, job := range jobs {
		jobsChannel <- job
	}

	close(jobsChannel)

	for result := range resultChannel {
		fmt.Printf("Received: %v\n", result)
	}

	fmt.Printf("It Took %s ms.\n", time.Since(startTime))
}
