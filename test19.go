//channel
package main

import (
	"fmt"
	//"time"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	msg := make(chan string, 1)
	msg1 := make(chan string, 1)

	go func() {
		ping(msg, "test1111")
		pong(msg, msg1)
		fmt.Println(<-msg)
	}()

	ping(msg, "test")
	pong(msg, msg1)
	fmt.Println(<-msg1)
}
