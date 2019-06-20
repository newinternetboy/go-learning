package main

import "fmt"

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//map
	var state = make(map[int]int)

	//状态锁
	var mutex = &sync.Mutex{}

	//计数
	var ops int64 = 0

	//100协程(读)
	for r := 0; r < 100; r++ {
		go func() {
			total := 0

			for {
				key := rand.Intn(5)
				//上锁
				mutex.Lock()
				//加值
				total += state[key]
				//释放锁
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				//手动让出资源控制权
				runtime.Gosched()
			}
		}()
	}

	//10个协程(写)
	for i := 0; i < 10; i++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			//增加状态计数
			atomic.AddInt64(&ops, 1)
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)

	//输出总次数
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	//输出最终的状态
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
