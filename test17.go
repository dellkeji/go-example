//通道学习

package main

import "fmt"

func main() {
	//创建channel,使用make(chan val-type)
	m_channel := make(chan string)
	m_channel_test := make(chan string)
	// 使用‘channel<-'发送一个新值到通道中
	go func() {
		m_channel <- "ping"
		m_channel_test <- "pong"

		fmt.Println("yyyy")
		fmt.Println(<-m_channel)
		fmt.Println(<-m_channel_test)
	}()

	go func() {
		fmt.Println("tttttttt")
	}()

	go test(10)
	//从channel中读取值`<-channel`
	msg := <-m_channel
	fmt.Println(msg)
}

func test(num int) {
	fmt.Println(num)
}
