//通道学习

package main

import "fmt"

func main() {
	//创建channel,使用make(chan val-type)
	m_channel := make(chan string)

	// 使用‘channel<-'发送一个新值到通道中
	go func() {
		m_channel <- "ping"
	}()

	//从channel中读取值`<-channel`
	msg := <-m_channel
	fmt.Println(msg)
}
