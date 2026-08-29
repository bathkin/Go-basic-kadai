package main

import "fmt"

func main() {
	n := 0
	/*
		for文の基本形
		① 初期化文: i := 1
			ループ開始時に一度だけ実行される。変数iを1で初期化。
		② 条件式: i <= 5
			ループ処理の前に毎回評価される。iが5以下ならループを続ける。
		③ 後処理文: i++
			{ }内の処理が終わるたびに実行される。iの値を1増やす。
	*/

	for {
		fmt.Println(n)
		if n == 3 {
			fmt.Println("終了処理実行")
			break
		}
		n++
	}
	fmt.Println("終了")

}
