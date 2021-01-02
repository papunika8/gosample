package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 5, 1, 3, 6, 42, 3, 856}
	numbers := bobbleSort(nums)
	fmt.Println(numbers)
}

func bobbleSort(numbers []int) []int {
	lenNums := len(numbers)
	for i := 0; i < lenNums; i++ {
		for j := 0; j < lenNums-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}
