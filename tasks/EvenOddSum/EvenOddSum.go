package main

import "fmt"

func evenOddSum(arr []int)  {
	var evenSum, oddSum int

	for _, value := range arr {
		if value % 2 == 0 {
			evenSum += value
		} else {
			oddSum += value
		}
	}

	fmt.Println("even: ", evenSum )
	fmt.Println("odd: ", oddSum)
}

func main() {
	arr := []int{1,2,3123,23,2342,3423,423,42,34,2,3434,23,5,534,5,5665,7,6,4,734}

	evenOddSum(arr)
}
