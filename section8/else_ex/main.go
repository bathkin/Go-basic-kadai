package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	num = 10
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	if num := add(5, 3); num <= 10 {
		fmt.Println(num, "は10以下の値")
	}
}
