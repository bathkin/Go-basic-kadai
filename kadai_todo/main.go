package main

import "fmt"

type ToDo struct {
	ID        int
	Title     string
	Completed bool
}

func (t *ToDo) Complete() {
	fmt.Printf("ID：%sのToDoを完了に更新します\n", t.Title)
	t.Completed = true
}
func printTodos(todos []ToDo) {
	for _, todo := range todos {
		status := "[未完了]"
		if todo.Completed {
			status = "[完了]"
		}
		fmt.Printf("%s（ID：%d）%s\n", status, todo.ID, todo.Title)
	}

}

func main() {
	todos := []ToDo{
		{ID: 1, Title: "学習計画", Completed: false},
		{ID: 2, Title: "環境構築", Completed: false},
		{ID: 3, Title: "基礎文法", Completed: false},
	}
	fmt.Println("--- 初期状態 ---")
	printTodos(todos)
	fmt.Println("\n--- 更新処理 ---")
	todos[0].Complete()
	todos[1].Complete()
	fmt.Println("\n--- 最終状態 ---")
	printTodos(todos)

}
