package main

import (
	"fmt"
	"mylib"
)

func main() {
	numbers := mylib.RandNums(1000, 100)
	numbers = quickSort(numbers)
	fmt.Println(numbers)
}

func quickSort(numbers []int) []int {
	quicksorting(numbers, 0, len(numbers)-1)
	return numbers
}

func quicksorting(numbers []int, low int, high int) []int {
	var partisionIndex int
	if low < high {
		partisionIndex = pertition(numbers, low, high)
		quicksorting(numbers, low, partisionIndex-1)
		quicksorting(numbers, partisionIndex+1, high)
	}
	return numbers
}

func pertition(numbers []int, low int, high int) int {
	i := low - 1
	pivot := numbers[high]
	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}
