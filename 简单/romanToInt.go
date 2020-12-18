package main

import "fmt"

// 罗马数字转整数
func romanToInt(s string) (r int) {
	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := getInt(s[i])
		if cur >= pre {
			r += cur
		} else {
			r -= cur
		}
		pre = cur
	}

	return

}

func getInt(r byte) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func main() {
	fmt.Println(romanToInt("IV"))
}
