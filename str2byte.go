package main

import "fmt"

// str2byte : tranform string foramt to byte format
func str2byte(str string) []byte {
	item := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		c := str[i]
		item[i] = c
	}
	return item
}

func main() {
	fmt.Println("String to byte...")
	str := "this is a test"
	fmt.Println(str2byte(str))
}
