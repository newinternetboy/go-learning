package main

import "fmt"
import "time"

func main() {
	msg := make(chan string)
	
	go func(){
		msg <- "ping"
	}()	
	message :=  <- msg
	fmt.Println(message)
	testCloseChannel()
	testMergeInput()
	time.Sleep(time.Second * 20)	
	many_goroutine_deal_channel()
}

/**测试通道关闭**/
func testCloseChannel() {
	//定义两个通道
	ch := make(chan int,5)
	sign := make(chan int, 2)
	
	//开启协程灌数据
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		
		//关闭通道
		close(ch)
		
		fmt.Println("the channel is closed!")

		sign <- 0
	}()

	//开启协程2来接受协程1灌入的数据，如果没有数据，跳出数据读取
	go func() {
		for {
			i, ok := <- ch
		
			fmt.Printf("%d, %v \n", i, ok)
	
			if !ok {
				break
			}
			time.Sleep(time.Second * 2)
		}
		sign <- 2
	}()
	
	fmt.Println(<- sign)
	fmt.Println(<- sign)
}

/** 多个输入同一个通道输出 **/
func testMergeInput() {
	input1 := make(chan int)
	input2 := make(chan int)
	output := make(chan int)
	for i := 0; i <=30; i++ {
	go func(in1, in2 <-chan int, out chan<- int) {
		select {
			case v := <- in1:
			out<- v
			case v := <- in2:
			out<- v	
		}	
	}(input1, input2, output)

	}
	
	//向input1通道灌数据
	go func() {
		for i := 0; i <= 19; i++ {
			input1 <- i
		}
	}()
	
	//向input2通道灌数据
	go func() {
		for i := 20; i <= 30 ; i++ {
			input2 <- i
		}
	}()

	//输出通道v的数据
	go func() {
		for {
			select {
				case value := <-output:
				fmt.Println("v通道数据:", value)
			}
		}
	}()
	
	//防止主线程提前退出
	fmt.Println("主线程退出")
}

/** 多个goroutine 同时获取同一个通道的值 **/
func many_goroutine_deal_channel(){
	chan1 := make(chan bool)
	chan1 <- true
	go func(chan1 <-chan bool){
		fmt.Println(<-chan1)
	}(chan1)

	go func(chan2 <-chan bool){
		fmt.Println(<-chan1)
	}(chan1)
}
