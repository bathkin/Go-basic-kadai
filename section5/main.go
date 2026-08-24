package main

import "fmt"

func main() {
	//型指定
	name := "Gopher"
	fmt.Println(name)
	name = "Go2"
	fmt.Println(name)
	//複数変数の指定
	var (
		country string = "Japan"
		city           = "Tokyo"
	)
	fmt.Println(country)
	fmt.Println(city)
	const pi = 3.14
	//pi = 3.2
	fmt.Println("円周率", pi)
}
