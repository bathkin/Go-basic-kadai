package main

import "fmt"

type User struct {
	Name  string
	Price int
}

func main() {
	u1 := User{
		Name:  "田中",
		Price: 100,
	}
	fmt.Println(u1.Name, u1.Price)
}
