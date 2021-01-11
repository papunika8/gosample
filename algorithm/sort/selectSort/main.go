package main

import (
	"fmt"
	"mylib"
)

func main() {
	numbers := []int{}
	numbers = mylib.RandNums(numbers, 1000000, 100000)
	numbers = selectSort(numbers)
	fmt.Print(numbers)
}

func selectSort(numbers []int) []int {
	lenNumbers := len(numbers)
	tmp := 0
	index := 0

	for i := 0; i < lenNumbers; i++ {
		flag := false
		tmp = numbers[i]
		for j := i + 1; j < lenNumbers; j++ {
			if tmp > numbers[j] {
				tmp = numbers[j]
				index = j
				flag = true
			}
		}
		if flag {
			numbers[i], numbers[index] = numbers[index], numbers[i]
		}
	}
	return numbers
}
