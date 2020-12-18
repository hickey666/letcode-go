package main

import "fmt"

// 找不同
func findTheDifference(s, t string) byte {
	sum := 0
	for _, ch := range s {
		sum -= int(ch)
	}
	for _, ch := range t {
		sum += int(ch)
	}
	return byte(sum)
}

func main() {
	res := findTheDifference("abcd", "abcde")
	fmt.Printf(string(res))
}
