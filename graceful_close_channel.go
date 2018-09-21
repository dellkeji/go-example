package main

import "fmt"

type T int

func isClose(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func main() {
	c := make(chan T)
	fmt.Println(isClose(c))
	close(c)
	fmt.Println(isClose(c))
}
