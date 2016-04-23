// array
package main

import "fmt"

func main() {
	fmt.Println("start here")

	var a [5]int
	fmt.Print(a)

	// initial the value
	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	var b [2][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			b[j][i] = i + j
		}
	}
	fmt.Println(b)
}
