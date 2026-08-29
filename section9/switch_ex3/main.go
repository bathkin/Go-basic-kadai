package main

import "fmt"

func main() {
	switch age := 18; {
	case age >= 60:
		fmt.Println("老人")
	case age >= 30:
		fmt.Printf("中年です")
	case age >= 10:
		fmt.Printf("青年です\n")
	default:
		fmt.Println("該当なし")
	}
}
