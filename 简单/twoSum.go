package main

import "fmt"

// 两数之和
func twoSum(nums []int, target int) []int{
	maps := make(map[int]int)
	for k, num1 := range nums {
		num2 := target - num1
		if _, ok := maps[num2]; ok {
			return []int{maps[num2], k}
		} else{
			maps[num1] = k
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
