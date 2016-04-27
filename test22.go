// timer

package main

import (
	"fmt"
	"time"
)

func main() {
	// create a timer, wait for seconds time
	timer1 := time.NewTimer(time.Second * 2)

	// block the next until the timer1 send message
	msg := <-timer1.C
	fmt.Println("time1 expired", msg)

	time.Sleep(time.Second * 2)
	fmt.Println("test")
}
