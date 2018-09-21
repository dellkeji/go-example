package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	log "github.com/cihub/seelog"
)

const (
	DateDefaultFormat = "2006-01-02"
	TimeFormat        = "15:04:05"
)

const (
	a = iota
	b
	c
	d
)

func main() {
	defer log.Flush()
	log.Info("%Date %Time test")
	kk := map[string]([]int){
		"a": []int{1, 4},
		"b": []int{6, 10},
	}
	fmt.Println(kk)
	fmt.Println(a, b, c)

	m := map[string]interface{}{
		"name":    "backy",
		"species": "dog",
	}
	mJson, _ := json.Marshal(m)
	fmt.Println(mJson)
	contentReader := bytes.NewReader(mJson)
	fmt.Println(string(mJson))
	d := json.NewDecoder(contentReader)
	var test111 map[string]interface{}
	d.Decode(&test111)
	fmt.Println(reflect.TypeOf(test111))
	fmt.Println(contentReader)

	switch v := reflect.ValueOf(3); v.Kind() {
	case reflect.Int:
		fmt.Println(111)
	default:
		fmt.Println(222)
	}

	test_val := []string{"a", "2", "3"}
	mm := []string{"4", "5"}
	dd := append(test_val, "5")

	s := []int{1, 2, 3}
	fmt.Println(s[:1])
	fmt.Println(dd)
	fmt.Println(mm)
}
