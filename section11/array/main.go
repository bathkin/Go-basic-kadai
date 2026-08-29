package main

import "fmt"

func main() {
	// 長さ3の配列を初期化
	scores := [3]int{80, 95, 72}
	fmt.Println("初期状態:", scores) // [80 95 72]

	// --- 要素の読み取り ---
	// インデックス 0 (先頭) の要素を読み取る
	length := len(scores)
	fmt.Println("長さ", length)
	firstScore := scores[0]
	fmt.Println("0番目の点数:", firstScore) // 80

	// インデックス 2 (3番目) の要素を読み取る
	thirdScore := scores[2]
	fmt.Println("2番目の点数:", thirdScore) // 72

	// --- 要素の書き込み (変更) ---
	// インデックス 1 (2番目) の値を 98 に変更する
	scores[1] = 70
	fmt.Println("変更後:", scores) // [80 98 72]
}
