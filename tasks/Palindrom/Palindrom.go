package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	var length int = len(str)

	for i := 0; i < length / 2; i++ {
		if str[i] != str[length - i - 1] {
			return false
		}
	}

	return true
}

func main() {
	var str string = "1234567890987654321"

	if isPalindrome(str) {
		fmt.Println("palindrom")
	} else {
		fmt.Println("palindrom emas")
	}
}
