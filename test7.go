// slice

package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s, s[1])

	fmt.Println(len(s))

	s = append(s, "d")

	s = append(s, "e", "f")

	fmt.Println(s)

	fmt.Println(s[1:2])
}
