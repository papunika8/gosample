package main

import (
	"fmt"
	"strings"
)

func main() {
	var moji string
	var kazu int
	fmt.Scan(&kazu, &moji)
	moji = atcoder(moji)
	fmt.Println(moji)

}

func atcoder(moji string) string {
	var result string
	var st string
	for len(moji) > 0 {
		st = string(moji[len(moji)-1:])
		result = string(st) + result
		moji = strings.Replace(moji, string(st), "", -1)
	}
	return result
}
