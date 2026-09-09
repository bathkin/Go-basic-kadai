package main

import "fmt"

type Order interface {
	~int | ~float64 | ~string
}

type Box[T any] struct {
	value T
}

func NewBox[T any](v T) *Box[T] {
	return &Box[T]{value: v}

}

func (a Box[T]) GetValue() T {
	return a.value
}

func Min[T Order](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("___ジェネリックなmin関数を実行")
	fmt.Printf("min(10, 20)= %v\n", Min(10, 20))
	fmt.Printf("min(3.14, 2.71)= %v\n", Min(3.14, 2.72))
	fmt.Printf("Min(\"Go\", \"Java\") = %v\n", Min("Go", "Java"))
	intBox := NewBox(123)

	v1 := intBox.GetValue()
	fmt.Printf("intBoxの中身：%v（型：%T）\n", v1, v1)
	strBox := NewBox("hello")
	v2 := strBox.GetValue()
	fmt.Printf("intBoxの中身：%v（型：%T）\n", v2, v2)
}
