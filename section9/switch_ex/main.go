package main

import "fmt"

func main() {
	score := 2

	switch score {
	case 1:
		fmt.Println("未達")
	case 2:
		fmt.Println("まずまず")
		fallthrough
	case 3:
		fmt.Println("素晴らしい")
	case 4:
		fmt.Println("評価対象外")
	default:
		fmt.Println("該当なし")
	}
}
