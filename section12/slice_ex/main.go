package main

import "fmt"

func main() {
	si := []int{10, 20, 30}
	fmt.Println("リテラル s1:", si)
	fmt.Println("リテラルの長さ ", len(si))
	se := make([]string, 3)
	fmt.Println("make se:", se)
	fmt.Println("リテラルの長さ ", len(se))
	si[1] = 40
	se[0] = "apple"
	se[1] = "banana"
	//se[2] = "snack"
	fmt.Println("se(代入後):", se)
}
