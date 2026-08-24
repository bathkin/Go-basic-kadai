package main

import "fmt"

func main() {
	var age int
	age = 30
	var count int = 10
	fmt.Println("年齢", age)
	fmt.Println("個数", count)
	//計算
	var sum = count + 5
	diff := count - 5
	prodct := count * 5

	fmt.Println("足し算", sum)
	fmt.Println("足し算", diff)
	fmt.Println("足し算", prodct)

	const pi float64 = 3.14
	radius := 5.0
	area := radius * radius * pi
	fmt.Println("面積", area)
	name := "go"
	fmt.Println(name + "学習中")
	var message string = `
	Goの学習
	6章:データ型
	`
	fmt.Println(message)
	isFinish := true
	fmt.Println("処理完了", isFinish)
	isLooding := false
	fmt.Println("ロード中", isLooding)
	result := (10 > 5)
	fmt.Println("10 > 5:", result)
	// 値を代入せずに宣言
	var i int
	var f float64
	var s string
	var b bool

	fmt.Println("intのゼロ値:", i)
	fmt.Println("float64のゼロ値:", f)
	fmt.Println("stringのゼロ値:", s)
	fmt.Println("boolのゼロ値:", b)
}
