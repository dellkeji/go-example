package main

import (
	"fmt"
	"os"
)

var limit = make(chan int, 3)

func main() {
	for _, w := range []int{1, 2, 3} {
		go func() {
			limit <- 1
			fmt.Println(w)
			<-limit
		}()
	}
	fmt.Println("test")
	select {}
	fmt.Println("test1")
	os.Exit(0)
}
