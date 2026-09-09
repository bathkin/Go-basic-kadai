package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("時刻取得")
	now := time.Now()

	fmt.Printf("現在の時刻%v\n", now)
	fmt.Printf("現在の時刻%v\n", now.Year())
	fmt.Printf("現在の時刻%v\n", now.Month().String())

	fmt.Printf("日：%d\n", now.Day())
	fmt.Printf("時：%d｜分：%d｜秒：%d\n", now.Hour(), now.Minute(), now.Second())

	const dataTime = "2006-01-02 15:04:05"
	formatData := now.Format(dataTime)
	fmt.Printf("日付＋時刻：%s\n", formatData)
	formatDataTime := now.Format("2006-01-02")
	fmt.Printf("日付：%s\n", formatDataTime)

	dataString := "2025-01-23 10:30:00"
	parse, err := time.Parse(dataTime, dataString)
	if err != nil {
		fmt.Println("パースに失敗しました：", err)
		return
	}
	fmt.Printf("パースした時刻: %v(型：%T)\n", parse, parse)
	twoHour := 2 * time.Hour
	twoTimwHour := now.Add(-twoHour)
	fmt.Printf("現在から2時間後：%v\n", twoTimwHour)
	oneYear := now.AddDate(-1, -1, 0)
	fmt.Printf("1年後：%v\n", oneYear)

	diff := twoTimwHour.Sub(now)
	fmt.Printf("2時間後と現在の差：%v\n", diff)

	targetTime := now.Add(24 * time.Hour) // 24時間後
	nowCopy := now
	fmt.Printf("比較元（now）：%v\n", now)
	fmt.Printf("比較対象（targetTime）：%v\n", targetTime)
	fmt.Printf("比較対象（nowCopy）：%v\n", nowCopy)
	if now.Before(targetTime) {
		fmt.Println("now は targetTime より「前（Before）」です。")
	}
	if targetTime.After(now) {
		fmt.Println("targetTime は now より「後（After）」です。")
	}
	// .Equal(t)：レシーバー（now）が t と同じ時刻か？
	if now.Equal(nowCopy) {
		fmt.Println("now と nowCopy は「同じ（Equal）」です。")
	}
}
