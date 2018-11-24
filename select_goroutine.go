package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("default")
		case <-cancel:
			fmt.Println("return")
			return
		}
	}
}

func main() {
	cancel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(&wg, cancel)
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
