package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := os.Open("/Users/takaaki/go-basic/section30/write/hello.txt")
	if err != nil {
		fmt.Println("読み込みに失敗", err)
		return
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

	//content := string(data)

	if err := scanner.Err(); err != nil {
		fmt.Println("ファイル読み込み中にエラーが発生しました：", err)
		return
	}
	fmt.Println("ファイルの読み込み完了しました")
}
