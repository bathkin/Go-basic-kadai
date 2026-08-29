package main

import (
	"errors"
	"fmt"
	"strconv"
)

func print(s string) {
	fmt.Printf("%sの変換\n", s)
	num, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("エラー", err)
		return
	}
	fmt.Println("成功", num)
}
func check(age int) (string, error) {
	if age > 20 || age < 10 {
		return "", errors.New("学生の年齢ではない")
	}
	message := fmt.Sprintf("学生：%d歳です\n", age)
	return message, nil

}
func main() {
	msg, err := check(-10)
	if err != nil {
		fmt.Println("失敗か", err)
	} else {
		fmt.Println("成功", msg)
	}
	print("123")
	print("abc")
}
