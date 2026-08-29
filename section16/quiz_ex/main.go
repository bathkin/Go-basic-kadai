package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	fmt.Printf("割られる数：%d｜割る数：%d\n", a, b)
	if b == 0 {
		return 0, errors.New("0で割ることはできません")
	}
	return a / b, nil
}
func main() {
	result, err := divide(1, 0)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	fmt.Println("計算結果:", result)
}
