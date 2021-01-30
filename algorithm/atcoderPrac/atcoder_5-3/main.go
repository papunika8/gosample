package main

import (
	"fmt"
	"math"
)

func main() {
	alfb := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	test := make(map[int]string, 26)
	for num, i := range alfb {
		test[num] = string(i)
	}

	var kazu int
	fmt.Scan(&kazu)
	//kazu = 123
	var kurai int
	var keta int
	for i := 1; ; i++ {
		j := math.Pow(36, float64(i))
		if int(j) > kazu {
			kurai = i
			break
		}
	}
	result := make([]string, kurai, kurai)

	for i := 0; ; i++ {
		if kazu < 36 {
			result[0] = test[kazu]
			break
		}
		t := math.Pow(36, float64(i))
		if kazu < int(t)*36 {
			keta = kazu / int(t)
			result[i] = test[keta]
			kazu = kazu - keta*int(t)
			i = 0
		}
	}

	var out string
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] == "" {
			out = out + "0"
			continue
		}
		out = out + result[i]
	}
	fmt.Println(out)
}
