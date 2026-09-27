package main

import (
	"fmt"
	"math/rand"
	"time"
)

func search(name string, result chan string) {
	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(delay)
	result <- fmt.Sprintf("Результат из %s (за %v)", name, delay)
}

func Z6() {
	sources := []string{"База A", "API B", "Кэш C", "Индекс D"}
	result := make(chan string, len(sources))

	go search(sources[0], result)
	go search(sources[1], result)
	go search(sources[2], result)
	go search(sources[3], result)

	first := <-result
	fmt.Println("Первый ответ:", first)
}
