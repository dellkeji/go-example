// time out
package main

import (
	"fmt"
	"time"
)

func main() {
	// create a channel and return the result for two seconds
	m_channel := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		m_channel <- "result 1"
	}()

	//这里使用一个‘select’实现超时，<- m_channel 等待结果
	// <- Time.After 等待超时时间1s后发送的值，由于select
	//默认处理第一个已准备好的接收操作，如果这个操作超过了允许的
	//1s的话，将会执行超时case
	select {
	case res := <-m_channel:
		fmt.Println(res)
	// set 3 to 2, print timeout
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}

	// not block by default
	select {
	case msg := <-m_channel:
		fmt.Println(msg)
	default:
		fmt.Println("this is a end")
	}

	close(m_channel)

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for num := range queue {
		fmt.Println(num)
	}
}
