package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// string is read only
// unsafe : 唯一不安全的是它们调用结果可能在后来的版本中返回不同的值
func byte2str(src []byte) (p string) {
	data := make([]byte, len(src))
	for i, c := range src {
		data[i] = c
	}
	hdr := (*reflect.Strimako_templatesngHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(src)

	return p
}

func main() {
	fmt.Println([]byte("the tat"))
	fmt.Println(byte2str([]byte("the tat")))
}
