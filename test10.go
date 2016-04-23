// 函数是go的核心

package main

import "fmt"

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	res, res1 := minus(3, 4)
	fmt.Println(res, res1)

	mult([]int{1, 2, 3}...)
}

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) (int, int) {
	return 3, 7
}

func mult(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}
