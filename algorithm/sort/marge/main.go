package main

import (
	"fmt"
	"mylib"
)

func main() {
	numbers := mylib.RandNums(1000, 100)
	nestNumbers := marge(numbers)
	fmt.Println(nestNumbers)

}

func marge(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	center := len(numbers) / 2
	left := numbers[:center]
	right := numbers[center:]

	left = marge(left)
	right = marge(right)

	i, j, k := 0, 0, 0
	tmp := make([]int, len(numbers))
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			tmp[k] = left[i]
			i++
		} else {
			tmp[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		tmp[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		tmp[k] = right[j]
		j++
		k++
	}
	return tmp
}
