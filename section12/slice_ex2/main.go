package main

import "fmt"

func main() {
	s2 := []int{40, 50}
	fmt.Println(s2)
	s2 = append(s2, 100, 120)
	fmt.Println(s2)
	s3 := s2[:3]
	fmt.Println(s3)
	s2[0] = 1100
	fmt.Println(s3)
	fmt.Println(s2)

}
