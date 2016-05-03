// worker pool

package main

import (
	"fmt"
	"time"
)

// 这是我们将要在多个并发实例中支持的任务了。这些执行者
// 将从 `jobs` 通道接收任务，并且通过 `results` 发送对应
// 的结果。我们将让每个任务间隔 1s 来模仿一个耗时的任务。
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "job id", j)
		time.Sleep(time.Second * 1)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	//设置3个worker，初始是阻塞的
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	//设置9个jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	//collect this result
	for a := 1; a <= 9; a++ {
		<-result
	}
}
