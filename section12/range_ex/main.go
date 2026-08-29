package main

import "fmt"

func main() {
	s2 := []string{"数学", "英語", "体育"}

	for i, v := range s2 {
		fmt.Println(i, v)
	}
	for _, v := range s2 {
		fmt.Println(v)
	}
	for i := range s2 {
		fmt.Println(i)
	}

}
