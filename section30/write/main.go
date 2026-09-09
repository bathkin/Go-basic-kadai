package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Println("ファイル読み込み失敗", err)
		return
	}
	defer file.Close()
	Lines := []string{
		"one",
		"two",
		"three",
	}
	fmt.Println("読み込み成功")
	for _, line := range Lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("書き込み失敗%v", err)
			return
		}
	}

}
