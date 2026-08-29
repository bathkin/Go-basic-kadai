package main

import "fmt"

func main() {
	food := map[string]int{
		"cake":  10,
		"bread": 30,
		"meat":  50,
	}
	food["cake"] = 60
	food["veg"] = 30
	delete(food, "cake")

	fmt.Println(food)
	fmt.Printf("パンの個数は%dです\n", food["bread"])
	fmt.Println(food["yamada"])
	value, ok := food["yamada"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("存在しません。")
	}
	for foods, count := range food {
		fmt.Printf("%sは%d個在庫がある\n", foods, count)
	}

}
