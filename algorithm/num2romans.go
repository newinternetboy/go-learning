package main
//遍历该key数组，找到当前数可以减去的最大值，查表返回其罗马数，拼接上减去该最大值后的数的罗马数，递归直至被减数与减数相等时返回结果。
import (
	"fmt"
	"sort"
)

var (
	table = map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
		4:    "IV",
		9:    "IX",
		40:   "XL",
		90:   "XC",
		400:  "CD",
		900:  "CM",
	}
	romans = func() []int {
		var keys []int
		for k := range table {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		return keys
	}()
)

func intToRoman(num int) string {
	subtrahend := romans[len(romans)-1]
	for i, v := range romans {
		if num < v {
			subtrahend = romans[i-1]
			break
		}
	}
	if 0 == num-subtrahend {
		return table[subtrahend]
	}
	return table[subtrahend] + intToRoman(num-subtrahend)
}

func main() {
	fmt.Println(intToRoman(519))
}
