package main

import (
	"fmt"
	//"sync/atomic"
	"time"
)

func main() {
	var cnt uint32 = 0
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 20; j++ {
				time.Sleep(time.Millisecond)
				cnt += 1
				//	atomic.AddUint32(&cnt, 1)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(cnt)
	//fmt.Println(atomic.LoadUint32(&cnt))
}
