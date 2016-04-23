//针对闭包实现匿名函数（不需要命名函数）
package main

import "fmt"

// 在函数中返回另一个函数
func intsq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// set val
	next_val := intsq()
	fmt.Println(next_val())
	fmt.Println(next_val())

	next_val_one := intsq()
	fmt.Println(next_val_one())

}
