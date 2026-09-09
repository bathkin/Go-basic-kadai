package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname,omitempty"`
}

func main() {
	fmt.Println("マーシャル")

	u1 := User{
		FirstName: "たかし", LastName: "ごうだ", Nickname: "ジャイアン",
	}
	u2 := User{
		FirstName: "ノビ", LastName: "のび太",
	}
	j1, _ := json.Marshal(u1)
	j2, _ := json.Marshal(u2)

	fmt.Println("マーシャルされたJSONデータ：")
	fmt.Println(string(j1))
	fmt.Println(string(j2))

	fmt.Println("\n--- 2. アンマーシャル（JSON -> Go）---")

	inputJson := []byte(`{"first_name":"二郎","last_name":"侍","nickname":"ジロー"}`)

	var u3 User
	err := json.Unmarshal(inputJson, &u3)
	if err != nil {
		fmt.Println("アンマーシャルに失敗しました：", err)
		return
	}
	fmt.Printf("%+v\n", u3)
}
