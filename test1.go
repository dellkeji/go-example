package main

import (
	"fmt"
	"strings"
)

const b = 5

func main() {
	//a := 5
	fmt.Println(b)
	fmt.Println("go" + "lang")

	// 整数和浮点数
	fmt.Println("1+1 = ", 2)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// 布尔
	fmt.Println(true && false)

	num := make(map[string]int)
	num["test"] = 1
	fmt.Println(num)
	for val := range num {
		fmt.Println(val)
	}
	test_str := "这是一个测试"
	counte_str := strings.Count(test_str, "") - 1
	fmt.Println(counte_str)

	l1 := len([]rune(test_str))
	fmt.Println(l1)
}

//package main
//
//import (
//	"fmt"
//	"log"
//
//	"gopkg.in/yaml.v2"
//)
//
//var data = `
//blog: xiaorui.cc
//best_authors: ["fengyun","lee","park"]
//desc:
//  counter: 521
//  plist: [3, 4]
//`
//
//// desc： subStruct
//type T struct {
//	Blog    string
//	Authors []string `yaml:"best_authors,flow"`
//	Desc    struct {
//		Counter int   `yaml:"Counter"`
//		Plist   []int `yaml:",flow"`
//	}
//}
//
//func main() {
//	t := T{}
//	//把yaml形式的字符串解析成struct类型
//	err := yaml.Unmarshal([]byte(data), &t)
//	//修改struct里面的记录
//	t.Blog = "this is Blog"
//	t.Authors = append(t.Authors, "myself")
//	t.Desc.Counter = 99
//	fmt.Printf("--- t:\n%v\n\n", t)
//	//转换成yaml字符串类型
//	d, err := yaml.Marshal(&t)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	fmt.Printf("--- t dump:\n%s\n\n", string(d))
//	for i := 0; i < 5; i++ {
//		var v int
//		fmt.Printf("%d ", v)
//		v = 5
//	}
//}
