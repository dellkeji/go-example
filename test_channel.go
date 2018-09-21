package main

import (
	"fmt"

	"time"
)

func main() {
	c := make(chan int) // an unbuffered channel
	go func() {
		fmt.Println(2)
		x := <-c // blocking here until a value is received.
		fmt.Println(3)
		c <- x * x // blocking here until the result is received.
	}()
	time.Sleep(1 * time.Second)
	c <- 3 // blocking here until the value is received.
	fmt.Println(1)
	time.Sleep(1 * time.Second)
	y := <-c       // blocking here until the result is received.
	fmt.Println(y) // 9
}
