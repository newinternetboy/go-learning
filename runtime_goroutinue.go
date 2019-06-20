package main
import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//runtime包中Gosched()函数会主动让出当前gorotinue的执行权限
	//下列两个goroutinue执行的顺序不分先后(试着用channel来让第一个先执行)
	wg.Add(2)
	channel1 := make(chan bool)
	go func(chann chan bool) {
		channel1 <- true
		//计数器减一
		for i := 0; i < 10; i++ {
			fmt.Printf("first %d\n", i)
			runtime.Gosched()
		}
		wg.Done()
	}(channel1) 
	
	go func(chann chan bool) {
		<- channel1
		for j := 10; j < 20; j++ {
			fmt.Printf("second %d\n", j)
			runtime.Gosched()
		}
		wg.Done()
	}(channel1)
	wg.Wait()
}
