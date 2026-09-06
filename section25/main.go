package main

import "fmt"

type Greeter interface {
	greet() string
}

type User struct{}

func (u User) greet() string {
	return "こんにちは"
}
func print(a any) {
	fmt.Printf("型：%T 値：%v,\n", a, a)
}
func main() {
	var g Greeter = User{}

	values := []any{
		123,
		"こんにちは",
		g,
	}
	for _, v := range values {
		print(v)
	}
	for _, v := range values {
		if num, ok := v.(int); ok {
			fmt.Printf("【int型】%dに10を加算します\n", num)
			num += 10
			continue
		}
	}
}
