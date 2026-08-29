package main

import "fmt"

func sayHello(name, job string, age int) int {
	result := age * 2
	fmt.Printf("%sは%d歳、職業は%sです\n", name, result, job)
	return age * 2
}

func main() {
	fmt.Println("main関数を開始します")
	sayHello("Alice", "食品メーカー", 25)
	sayHello("Bob", "エンジニア", 30)
	fmt.Println("main関数を終了します")
}
