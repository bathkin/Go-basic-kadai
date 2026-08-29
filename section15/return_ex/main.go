package main

import "fmt"

func info() (int, string) {
	return 30, "田中"
}

func main() {
	age, _ := info()
	fmt.Printf("氏名:年齢:%d\n", age)
}
