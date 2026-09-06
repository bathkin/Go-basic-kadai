package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (t Rectangle) Area() float64 {

	return t.width * t.height
}

type Triangle struct {
	base   float64
	height float64
}

func (r Triangle) Area() float64 {
	return (r.base * r.height) / 2
}
func printArea(a Shape) {
	fmt.Printf("この図形の面積は%0.1fです\n", a.Area())
}
func main() {
	shapes := []Shape{
		Rectangle{width: 10, height: 5},
		Triangle{base: 20, height: 4},
		Rectangle{width: 7, height: 2},
	}
	totalArea := 0.0
	for _, s := range shapes {
		fmt.Printf("%T計算します%0f\n", s, s.Area())
		totalArea += s.Area()
	}
	fmt.Printf("--- すべての図形の合計面積：%f ---\n", totalArea)
}
