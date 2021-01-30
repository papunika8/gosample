package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	var tmp rune
	count := 1

	for finish, i := range a {

		if i == tmp {
			count++
		} else {
			tmp = i
			count = 1
		}
		if count == 3 {
			fmt.Println(string(i))
			break
		}
		if finish == 4 {
			fmt.Println("draw")
		}
	}
}
