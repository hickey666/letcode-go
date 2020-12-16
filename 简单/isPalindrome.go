package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x > 0) {
		return false
	}
	y := 0
	z := x
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	return y == z
}

func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		x = x / 10
	}
	return y == x || x == y/10
}

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(isPalindrome(num))
}
