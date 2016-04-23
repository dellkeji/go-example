//指针

package main

import "fmt"

func zeroval(n int) {
	n = 0
}

func zeroptr(n *int) {
	*n = 0
}

func main() {
	i := 1
	fmt.Println(i)

	zeroval(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)

	fmt.Println(&i)
}
