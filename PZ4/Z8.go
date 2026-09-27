package main

import (
	"fmt"
	"sync"
	"time"
)

func Z8() {
	jobs := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Воркер %d обработал запрос %d\n", id, job)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
