package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name  string `json:"product_name"`
	Price int    `json:"price"`
}

func main() {
	goods := []Product{
		{Name: "冷蔵庫", Price: 15000},
		{Name: "テレビ", Price: 30000},
	}
	jsonData, err := json.Marshal(goods)
	if err != nil {
		fmt.Println("失敗")
		return
	}

	fmt.Println(string(jsonData))

}
