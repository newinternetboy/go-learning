package main

import "fmt"

func main(){
	var a int = 10
	var b int = 20
	var ret int
	
	/* 调用函数并返回最大值 */
	ret = max(a, b)
	fmt.Printf("最大值是:%d\n",ret)
	
	/* 交换两个数值 */
	a, b = swap(a, b)
	fmt.Printf("交换后的数值a=%d，b=%d\n",a, b)
	
	/* 闭包 */
	nextNumber := getSequence();
	fmt.Printf("%d\n",nextNumber())
	fmt.Printf("%d\n",nextNumber())
	
	/* 返回多个值 */
	x, y := returnMultiValue()
	fmt.Println(x,y)
	
	/* 变参函数 */
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	/* 匿名闭包函数 */
	incr := incr()
	fmt.Println(incr())
	fmt.Println(incr())
	fmt.Println(incr())

	/* 递归 */
	fmt.Println(fact(4))
}

func max(num1 int, num2 int) int {
	var result int
	if num1 > num2{
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(num1, num2 int) (int, int){
	return num2, num1
}

/* 闭包函数 */
func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}

/* 返回多个值 */
func returnMultiValue() (int, int) {
	return 1, 2
}


/* 可变参数函数 */
func sum(nums ...int) {
	fmt.Print(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

/* 匿名函数 闭包*/
//闭包里面的参数只有在脚本执行完毕之后才会释放
func incr() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/* 递归 */
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
