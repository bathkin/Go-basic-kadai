package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type UserDefault struct {
	Id   int
	Name string
}
type UserCustom struct {
	Id   int
	Name string
}

func (u UserCustom) String() string {
	return fmt.Sprintf("【ID：%d】【Name：%s】\n", u.Id, u.Name)
}

func runStringer() {
	defUser := UserDefault{Id: 1, Name: "デフフォルト値"}
	custom := UserCustom{Id: 5, Name: "カスタム値"}
	fmt.Println(defUser)
	fmt.Println(custom)
}

func runReaderwriter() {
	fmt.Println("\n--- io.Copy()でデータをコピー ---")
	sourcerReader := strings.NewReader("データコピーテスト")
	var destinationBuffer bytes.Buffer
	fmt.Println("コピー前のdst：", destinationBuffer.String())
	written, _ := io.Copy(&destinationBuffer, sourcerReader)
	fmt.Printf("\n転送完了。バイト数: %d\n", written)
	fmt.Println("コピー後のdst：", destinationBuffer.String())
}
func main() {
	runStringer()
	runReaderwriter()
}
