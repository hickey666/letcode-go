package main

import "fmt"

func removeElement(nums []int, val int) int {
	var i int
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	nums := [4]int{3,2,2,3}
	fmt.Println(removeElement(nums[0:], 3))
}
