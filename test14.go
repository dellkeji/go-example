// struct
package main

import "fmt"

func main() {
	fmt.Println(person{"bellke", 22})

	person_info := person{name: "bellke", age: 30}
	fmt.Println(person_info)

	fmt.Println(person_info.name)

	person_info.age = 31
	fmt.Println(person_info)

	person_info_ptr := &person_info
	fmt.Println(person_info_ptr.name)
}

type person struct {
	name string
	age  int
}
