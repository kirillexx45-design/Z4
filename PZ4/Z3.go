package main

import (
	"fmt"
	"sync"
	"time"
)

func Z3() {
	const (
		numRequests = 15
		rateLimit   = 200 * time.Millisecond
	)
	tick := time.Tick(rateLimit)
	var wg sync.WaitGroup
	for i := 1; i <= numRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-tick
			fmt.Printf("Запрос %d обработан в %s\n", id, time.Now().Format("15:04:05.000"))
		}(i)
	}

	wg.Wait()
	fmt.Println("Все запросы обработаны")
}
