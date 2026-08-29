package main

import "fmt"

func main() {

	for n := 1; n <= 9; n++ {
		for m := 1; m <= 9; m++ {
			fmt.Printf("%d ×　%d = %2d\n", n, m, n*m)
		}
	}
}
