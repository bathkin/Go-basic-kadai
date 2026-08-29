package main

import "fmt"

func main() {
	japanCity := map[string]string{
		"hokkaido": "sapporo",
		"hyogo":    "kode",
		"tokyo":    "渋谷",
	}
	fmt.Printf("東京の県庁所在地は%sです'\n", japanCity["tokyo"])
}
