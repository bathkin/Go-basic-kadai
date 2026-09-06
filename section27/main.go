package main

import "fmt"

const (
	errorCodeNagative = 1 //負の数のエラー
	errorCodeZero     = 2 //ゼロ除数
)

type ErrorCode struct {
	Code    int
	Message string
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("error %d,%s\n", e.Code, e.Message)
}

func Divade(A, B int) (float64, error) {
	if A < 0 || B < 0 {
		return 0, &ErrorCode{
			Code:    errorCodeNagative,
			Message: "負の数は非対称です",
		}
	}
	if B == 0 {
		return 0, &ErrorCode{
			Code:    errorCodeZero,
			Message: "ゼロは割れない数です",
		}
	}
	return float64(A) / float64(B), nil
}

func Run(a, b int) {
	fmt.Printf("--- %d / %d を計算 ---\n", a, b)
	result, err := Divade(a, b)
	if err != nil {
		if e, ok := err.(*ErrorCode); ok {
			fmt.Println("独自エラー（AppError）が発生しました。")
			fmt.Printf("  %s\n", e.Error())
			switch e.Code {
			case errorCodeNagative:
				fmt.Println("  対処方法：正の数を入力してください。")
			case errorCodeZero:
				fmt.Println("  対処方法：0以外の値を入力してください。")
			default:
				fmt.Println("対処方法：担当者に連絡してください。")
			}
		} else {
			fmt.Printf("  標準エラーが発生しました：%v\n", err)
		}
	} else {
		fmt.Printf("計算結果：%f\n", result)
	}
}

func main() {
	Run(9, 2)
	Run(-5, 2)
	Run(10, 0)
}
