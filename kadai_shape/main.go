package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Message string
	Field   string
	Value   float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("[バリデーションエラー] %s：%s（Value：%.2f）", e.Field, e.Message, e.Value)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}
func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * c.Radius * 3.14
}
func validateShape(s Shape) error {
	switch v := s.(type) {
	case Rectangle:
		if v.Width <= 0 {
			return &ValidationError{
				Message: "幅は正の数を指定してください。",
				Field:   "Rectangle.Width",
				Value:   v.Width,
			}
		}
		if v.Height <= 0 {
			return &ValidationError{
				Message: "高さは正の数を指定してください。",
				Field:   "Rectangle.Height",
				Value:   v.Height,
			}
		}
	case Circle:
		if v.Radius <= 0 {
			return &ValidationError{
				Message: "半径は正の数を指定してください。",
				Field:   "Circle.Radius",
				Value:   v.Radius,
			}
		}
	default:
		return fmt.Errorf("測定不能です")
	}
	return nil

}

func main() {
	// 処理対象の図形スライスを作成
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0}, // 不正（幅が不正）
		Circle{Radius: -1.0},                // 不正（半径が不正）
		Rectangle{Width: 1.0, Height: 0.0},  // 不正（高さが不正）
	}
	calcShapes := []Shape{}
	for i, s := range shapes {
		err := validateShape(s)
		if err == nil {
			fmt.Printf("%d : %T => バリデーションOK\n", i+1, s)
			calcShapes = append(calcShapes, s)
		} else {
			var ve *ValidationError
			if errors.As(err, &ve) {
				fmt.Printf("#%d：%T => 独自エラーが発生しました\n", i+1, s)
				fmt.Printf("  [詳細] Field=%s, Value=%.2f, Message=%s\n", ve.Field, ve.Value, ve.Message)
			} else {
				fmt.Printf("#%d：%T => 標準エラーが発生しました：%v\n", i+1, s, err)
			}
		}
	}
	for _, s := range calcShapes {
		fmt.Printf("図形：%T｜面積：%.2f ｜周長：%.2f\n", s, s.Area(), s.Perimeter())
	}

}
