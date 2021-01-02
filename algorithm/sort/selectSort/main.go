package main

import "fmt"

func main() {
	num := []int{2, 19, 4, 8, 5, 6, 99, 3, 10}
	numbers := selectSort(num)
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
