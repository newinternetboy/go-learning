package main

import "fmt"

func main() {
	nums := []int{0, 4, 3, 0}
	target := 0
	keys := twoSum(nums, target)
	fmt.Println(keys)
	fmt.Println(len(keys))
}
//返回数组中两个元素之和等于目标值的下标
func twoSum(nums []int, target int) []int {
    keys := make([]int,2)
    for k1, v1 := range nums {
            for k2, v2 := range nums {
                if k1 != k2 && v1 + v2 == target {
                    keys[0] = k1
                    keys[1] = k2
                }
            }
    }
    return keys
}


