package main

import (
	"errors"
	"fmt"
	"strconv"
)

func calcPrice(ageStr string) (int, error) {
	if ageStr == "" {
		return 0, errors.New("年齢が入力されていません")
	}
	age, err := strconv.Atoi(ageStr)

	if err != nil {
		return 0, fmt.Errorf("年齢の変換に失敗しました: %w", err)
	}
	if age < 0 {
		// %wを使わない例：独自のバリデーションエラー
		return 0, fmt.Errorf("年齢が負の値です：%d", age)
	}
	var price int
	switch {
	case age >= 65:
		price = 1500
	case age >= 20:
		price = 3000
	default:
		price = 1800
	}
	return price, nil
}

func main() {
	inputs := []string{"25", "abc", "10", "-5", ""}
	for _, input := range inputs {
		fmt.Printf("--- 入力：%qの場合 ---\n", input)
		price, err := calcPrice(input)
		if err != nil {
			fmt.Printf("エラー発生: %v\n", err)

			if errors.Is(err, strconv.ErrSyntax) {
				fmt.Println("[Is] 判定：入力が数値の形式ではありませんでした")
			}
			var numErr *strconv.NumError
			if errors.As(err, &numErr) {
				fmt.Printf("[As] 判定：変換対象%qが不正です（詳細：%v）\n", numErr.Num, numErr.Err)
			}
		} else {
			fmt.Printf("料金：%d円\n", price)
		}
	}
}
