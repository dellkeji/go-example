package main

import (
	"fmt"
)

type test struct {
	name string
}

func (c test) Test() string {
	return c.name
}

func (c *test) Test1() error {
	c.name = "test112"
	return nil
}

func main() {
	c := test{name: "test111"}
	fmt.Println(c.Test())
	c.Test1()
	fmt.Println(c.name)

}
