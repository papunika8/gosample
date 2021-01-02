package main

import (
	"fmt"
	"os"
	"errors"
	"log"
)

var foo, bar string = "fooooo", "baaaar"

var small, large int = 10, 20

func swap(i, j int) (int, int) {
	return j, i
}

func main() {
	x, y := 3, 4
	x, y = swap(x, y)
	//fmt.Println(x, y) // 4, 3

	//x = swap(x, y) // コンパイルエラー

	x, _ = swap(x, y) // 第二戻り値を無視
	fmt.Println(x)    // 3

	//swap(x, y) // コンパイル，実行ともに可能
	foobar()
	moshimo(small, large)
	roop()

	n, err := div(10, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	nintendo()
	sum(2, 4)
	arr := [...]string{"a", "b", "c", "d"}
	fn(arr)
	fmt.Println(arr)
	text := []string{"t", "e", "s", "t"}
	addtxt(text)
	maps()
	var i int = 10
	callByValue(i)
	fmt.Println(i)
	callbyRef(&i)
	fmt.Println(i)
}

func foobar() {
	fmt.Println(foo, bar)
}

func moshimo(a, b int) {
	if a > b {
		fmt.Println("a is larger than b")
	} else if a < b {
		fmt.Println("b is larger than a")
	} else {
		fmt.Println("a equals b")
	}
}

func roop() {
	i := 0
	for {
		i++
		if i > 10 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Print(i)
	}
}

func nintendo() {
	for i := 1; i <= 15; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		default:
			fmt.Println(i)
		}
	}
}

func div(i, j int) (result int, err error) {
	if j == 0 {
		err = errors.New("divied by zero")
		return
	}
	result = i / j
	return
}

var sum func(v, m int) = func(v, m int) {
	fmt.Println(v + m)
}

func fn(arr [4]string) {
	arr[0] = "x"
	fmt.Println(arr)
}

func addtxt(s []string) {
	s = append(s ,".txt")
	for _, u := range s {
		fmt.Println(u)
	} 
}

func maps() {
	month := map[int]string{
		1: "January",
		2: "February",
	}
	_, ok := month[1]
	if ok {
		fmt.Println(month[1])
	}else {
		fmt.Println(month)
	}
}

func callByValue(i int) {
	i = 20
}

func callbyRef(i *int) {
	*i = 20
}

func filer() {
	file, err := os.Open("./error.txt")
	if err != nil {
		fmt.Println("Faild")
		panic(errors.New("miss"))
	}
	fmt.Println("Succses!")
	defer file.Close()
}
