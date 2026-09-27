package main

import (
	"fmt"
	"sync"
)

func Z2() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				results <- job * job
			}
		}()
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
