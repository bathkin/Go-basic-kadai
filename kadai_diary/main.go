package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	TimeStamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

const fileName = "diary.json"

func loadDiary() ([]Entry, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}
		return nil, fmt.Errorf("読み込みに失敗しました: %w", err)
	}
	var entries []Entry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, fmt.Errorf("JSON変換に失敗しました: %w", err)
	}

	return entries, nil
}

func saveDiary(entries []Entry) error {
	data, err := json.Marshal(entries)
	if err != nil {
		return fmt.Errorf("書き込みに失敗しました：%w", err)
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("書き込みに失敗しました：%w", err)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使い方：go run main.go [add|list] ...")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		// add の場合はタイトル(os.Args[2])と本文(os.Args[3])が必要なため引数が4つ以上あるかチェック
		if len(os.Args) < 4 {
			fmt.Println("使い方: go run main.go add \"タイトル\" \"本文\"")
			return
		}

		title := os.Args[2]
		content := os.Args[3]

		// 既存の日記を読み込む
		entries, err := loadDiary()
		if err != nil {
			fmt.Println("エラー:", err)
			return
		}

		// 新しい日記を追加して保存
		newEntry := Entry{
			TimeStamp: time.Now(),
			Title:     title,
			Content:   content,
		}
		entries = append(entries, newEntry)

		if err := saveDiary(entries); err != nil {
			fmt.Println("保存エラー:", err)
			return
		}

		fmt.Println("日記を記録しました！")

	case "list":
		entries, err := loadDiary()
		if err != nil {
			fmt.Println("エラー:", err)
			return
		}

		if len(entries) == 0 {
			fmt.Println("日記はまだありません。")
			return
		}
		for _, entry := range entries {
			dateStr := entry.TimeStamp.Format("2006-01-02 15:04:05")

			fmt.Println("--- 日記データ ---")
			fmt.Printf("日時：%s\n", dateStr)
			fmt.Printf("タイトル：%s\n", entry.Title)
			fmt.Println("本文：")
			fmt.Println(entry.Content)
		}

	default:
		fmt.Println("不明なコマンドです：", command)
	}
}
