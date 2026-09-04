package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "tst"
	s2 := " "
	s3 := "World"
	result := s1 + s2 + s3
	fmt.Println(result)
	words := []string{"Go", "is", "awesome"}
	wordsSpilit := "Go , is , awesome"

	resultJoin := strings.Join(words, ",")
	resultSplit := strings.Split(wordsSpilit, ",")
	resultContain := strings.Contains(wordsSpilit, "Go")
	fmt.Println(resultSplit[0], ",")
	fmt.Println(resultJoin)
	fmt.Println(resultContain)
}
