// for 是go中唯一的循环结构
package main

import "fmt"

func main() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带任何条件的循环，主要是break
	for {
		fmt.Println("loop")
		break
	}
}
