package main

import (
	"fmt"
)

type test struct {
	x, y int
}

type test1 struct {
	test
	mm int
}

// Squal :
func (t *test) Squal() int {
	return t.x * t.y
}

func (t test) Squal1() int {
	return t.x * t.y
}

func main() {
	fmt.Println("start")
	t := test{
		x: 1,
		y: 5,
	}
	// GO compiler support flow default
	fmt.Println(t.Squal())
	fmt.Println((&t).Squal())

	t1 := test1{
		test: test{
			x: 1,
			y: 5,
		},
		mm: 5,
	}
	fmt.Println(t1.x)
	a := func() int { return 1 * 5 }()
	fmt.Println(a)
	fmt.Println(test.Squal1(t))
}
