package main

import "fmt"

func main() {
	// 声明变量，可以一个或者多个
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// go 可以自己猜测初始化的变量类型
	var d = true
	fmt.Println(d)

	// int 默认为0
	var e int
	fmt.Println(e)

	// 声明变量的简写
	f := "short"
	fmt.Println(f)
}
