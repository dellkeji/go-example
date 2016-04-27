// 打点器_ 则是当你想要在固定的时间间隔重复执行
// 可以和定时器样停止
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("ticket", t)
		}
	}()

	//在指定的时间停止
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("tttt")
}
