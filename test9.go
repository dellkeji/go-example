package main

import "fmt"

func main() {
	nums := []int{3, 2, 1}
	for i, num := range nums {
		fmt.Println(i, num)
	}
}
