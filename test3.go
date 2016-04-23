package main

import "fmt"
import "math"

// `const` 用于声明一个变量
const s string = "const"

func main() {
	fmt.Println(s)

	// `const` 语句可以出现在任何`var`语句的地方
	const n = 500000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(n))
}
