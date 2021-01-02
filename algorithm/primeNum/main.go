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
	//素数を格納するスライスを定義する。この時スライスの中身で割れたら素数と判定するため先に2を入れておく。
	result := []int{2}
	// 時間の計測開始
	start := time.Now()
	//3 > i > n の数字をiに格納する。
	for i := 3; i <= *n; i++ {
		//falgをtrueで初期化する。
		flag := true
		//iをスライスの中身で繰り返し割り算する
		for _, v := range result {
			//vの2乗がiを上回ったらループを抜ける
			if v*v > i {
				break
			}
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

	//1を一番最初に追加する
	app := []int{1}
	result = append(app, result...)
	//スライスの中身を出力する
	for _, v := range result {
		fmt.Printf("%d ", v)
	}

	// 時間の計測終了
	end := time.Now()
	fmt.Printf("\n%f秒\n", (end.Sub(start)).Seconds())
}
