//默认通道是 无缓冲的，这意味着只有在对应的接受（`<- chan`）
// 通道准备好时，才允许进行发送（`chan <-`）,允许在没有对应
// 接收方的情况下，缓存限定数量的值

package main

import "fmt"

func main() {
	// 允许缓冲两个值
	m_chan := make(chan string, 2)

	m_chan <- "buff"
	m_chan <- "chan"

	fmt.Println(<-m_chan, <-m_chan)
}
