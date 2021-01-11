package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 4, 7, 8, 10, 15, 20, 25}
	num1 := binaryRecursion(numbers, 9)
	fmt.Println(num1)
	num2 := binarySearch(numbers, 9)
	fmt.Println(num2)
}

func binarySearch(numbers []int, num int) (indexNum int) {
	left, right := 0, len(numbers)-1

	for left <= right {
		midol := (left + right) / 2
		if numbers[midol] == num {
			return midol
		} else if numbers[midol] < num {
			left = midol + 1
		} else {
			right = midol - 1
		}
	}
	return -1
}

func binaryRecursion(numbers []int, num int) (indexNum int) {
	middol := search(numbers, num, 0, len(numbers)-1)
	return middol
}

func search(numbers []int, num int, left int, right int) (indexNum int) {
	if left <= right {
		return -1
	}
	midol := (left + right) / 2
	if numbers[midol] == num {
		return midol
	} else if numbers[midol] < num {
		return search(numbers, num, midol+1, right)
	} else {
		return search(numbers, num, left, midol+1)
	}

}
