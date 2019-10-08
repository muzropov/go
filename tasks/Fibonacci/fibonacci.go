package main

import "fmt"

//simple fibonacci
func fibonacci1(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci1(n - 1) + fibonacci1(n - 2)
	}
}

//fibonacci which starting with x, y
func fibonacci2(x, y, n int) int {
	if n == 1 {
		return x
	} else if n == 2 {
		return y
	} else {
		return fibonacci2(y, x + y, n - 1)
	}
}

func main() {
	fmt.Println(fibonacci1(10)) // 55
	fmt.Println(fibonacci2(1, 2, 10)) // 89
}
