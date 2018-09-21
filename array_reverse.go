package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	test := []int{1, 2, 3, 4}
	reverse(test)
	fmt.Println(test)
}
