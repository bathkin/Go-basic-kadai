package main

import "fmt"

type Car struct {
	name  string
	speed int
}

func (c Car) Run() {
	if c.speed > 0 {
		fmt.Printf("%sが%dkmで走っています\n", c.name, c.speed)
	} else {
		fmt.Printf("%sは停車中です\n", c.name)
	}
}
func (c Car) AcceleA() {
	c.speed += 10
}

func (c *Car) AcceleB() {
	c.speed += 10
}

func main() {
	taxi := Car{
		name:  "タクシー",
		speed: 0,
	}
	truck := Car{
		name:  "トラック",
		speed: 0,
	}
	taxi.Run()
	truck.Run()

	taxi.AcceleA()
	truck.AcceleB()
	taxi.Run()
	truck.Run()

}
