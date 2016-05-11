package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// create map
	var state = make(map[int]int)
	fmt.Println(state)
	// init mutex
	var mutex = &sync.Mutex{}

	var ops int64 = 0

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				fmt.Println(state[key], key)
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				// 为了确保这个 Go 协程不会再调度中饿死，我们
				// 在每次操作后明确的使用 `runtime.Gosched()`
				// 进行释放。这个释放一般是自动处理的，像例如
				// 每个通道操作后或者 `time.Sleep` 的阻塞调用后
				// 相似，但是在这个例子中我们需要手动的处理。
				runtime.Gosched()
			}
		}()
	}
	// write the state
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	// operate the state and mutex during 1s
	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
