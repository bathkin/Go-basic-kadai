package main

import (
	"fmt"
)

func main() {
	fmt.Println("テスト")
	version := 1.22
	language := "Go"
	fmt.Println(language)
	fmt.Println(language, "は", version, "です")
	fmt.Printf("%Tのバージョンは\n%.2fです\n", language, version)
	//演習
	goods := "パソコン"
	value := 50000
	review := 3.7
	fmt.Printf("%vの価格は%d円で、レビュー評価は%.1fです。\n", goods, value, review)
}
