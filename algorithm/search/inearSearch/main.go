package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 4, 7, 8, 10, 15, 20, 25}
	num := inearSearch(numbers, 15)
	fmt.Println(num)
}

func inearSearch(numbers []int, num int) (indexNum int) {
	for i := 0; i < len(numbers); i++ {
		if num == numbers[i] {
			return i
		}
	}
	return -1
}
