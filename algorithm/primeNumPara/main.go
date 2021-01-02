package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//引数で受け取る(デフォルトで100までの素数を出力)
	var n = flag.Int("num", 100, "")
	flag.Parse()
	quarter := *n / 100
	minnum1, maxnum1 := 3, quarter*38
	minnum2, maxnum2 := quarter*38, quarter*62
	minnum3, maxnum3 := quarter*62, quarter*82
	minnum4, maxnum4 := quarter*82, *n+1
	var result1, result2, result3, result4 []int
	ch1, ch2, ch3, ch4 := make(chan bool), make(chan bool), make(chan bool), make(chan bool)
	// 時間の計測開始
	start := time.Now()

	//素数判定
	go func() {
		result1 = primeNumber(result1, minnum1, maxnum1)
		ch1 <- true
	}()
	go func() {
		result2 = primeNumber(result2, minnum2, maxnum2)
		ch2 <- true
	}()
	go func() {
		result3 = primeNumber(result3, minnum3, maxnum3)
		ch3 <- true
	}()
	go func() {
		result4 = primeNumber(result4, minnum4, maxnum4)
		ch4 <- true
	}()

	<-ch1
	<-ch2
	<-ch3
	<-ch4

	//1を一番最初に追加する
	app := []int{1, 2}
	result1 = append(app, result1...)
	//スライスの中身を出力する
	for _, v := range result1 {
		fmt.Printf("%d ", v)
	}
	for _, v := range result2 {
		fmt.Printf("%d ", v)
	}
	for _, v := range result3 {
		fmt.Printf("%d ", v)
	}
	for _, v := range result4 {
		fmt.Printf("%d ", v)
	}
	// 時間の計測終了
	end := time.Now()
	fmt.Printf("\n%f秒\n", (end.Sub(start)).Seconds())
}

func primeNumber(result []int, minnum int, maxnum int) []int {
	//minnum > i > maxnum の数字をiに格納する。
	for i := minnum; i < maxnum; i++ {
		//falgをtrueで初期化する。
		flag := true
		//iをスライスの中身で繰り返し割り算する
		for v := 2; ; v++ {
			//vの2乗がiを上回ったらループを抜ける
			if v*v > i { //この行を追記
				break //この行を追記
			} //この行を追記
			//割り切れたらflagをfalseにしてループを抜ける
			if i%v == 0 {
				flag = false
				break
			}
		}
		//すべて割り切れなければ、素数のスライスに追加する。
		if flag == true {
			result = append(result, i)
		}
	}
	return result
}
