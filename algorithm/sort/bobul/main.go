package main

import (
	"fmt"
	"mylib"
)

func main() {
	numbers := mylib.RandNums(1000, 100)
	numbers = bobbleSort(numbers)
	fmt.Println(numbers)
	fmt.Println(cap(numbers), len(numbers))
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
