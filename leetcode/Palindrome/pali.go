package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	j := len(str) - 1
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(1001))
	fmt.Println(isPalindrome(100))
}
