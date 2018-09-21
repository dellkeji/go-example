package main

import "fmt"

func main() {
	demo := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("test", t)
		case bool:
			fmt.Println("test1", t)
		default:
			fmt.Println("other", t)
		}
	}
	demo(1)
	demo(true)
}
