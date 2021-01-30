package main

import (
	"fmt"
)

func insertion(numbers []int) []int {
	lenNums := len(numbers)
	for i := 1; i < lenNums; i++ {
		for j := i - 1; j >= 0; j-- {
			if numbers[i] < numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
				i = -1
				j = -1
			}
		}
	}
	return numbers
}

func main() {
	nums := []int{0021, 4, 006, 343, 7, 3, 35, 98}
	numbers := insertion(nums)
	fmt.Println(numbers)
}
