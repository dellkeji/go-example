// struct function

package main

import "fmt"

type test interface {
	area() float32
	perim() float32
}

// 实现对应接口中的方法
type area struct {
	a, b int
}

func (a area) area() int {
	return a.a * a.b
}

func (a area) perim() int {
	return (a.a + a.b) * 2
}

func main() {
	a := area{10, 11}
	fmt.Println(a.area())
}
