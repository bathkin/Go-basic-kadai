package main

import "fmt"

func pass(val int) {
	val = 555
	fmt.Printf("%d\n", val)
}

func passBy(ptr *int) {
	*ptr = 119
	fmt.Printf("%d\n", ptr)
}

func main() {
	x1 := 123
	x2 := "こんにちは"

	var p1 *int
	p1 = &x1
	p2 := &x2

	fmt.Printf("%v, 値%d\n", p1, *p1)
	fmt.Printf("%v, 値%s\n", p2, *p2)
	*p1 = 456
	fmt.Printf("%v, 値%d\n", p1, *p1)
	num := 1
	pNum := &num
	fmt.Printf("[before] num：%d\n", num)
	pass(num)
	fmt.Println(num)
	fmt.Printf("[before] num：%d\n", num)
	passBy(pNum)
	fmt.Println(num)
}
