package main

import (
	"fmt"
	"net/http"
	"sync"
)

func Z4() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://reddit.com",
		"https://wikipedia.org",
	}

	jobs := make(chan string)
	results := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				resp, err := http.Get(url)
				if err != nil {
					results <- fmt.Sprintf("%s -> ошибка", url)
					continue
				}
				resp.Body.Close()
				results <- fmt.Sprintf("%s -> %s", url, resp.Status)
			}
		}()
	}

	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
}
