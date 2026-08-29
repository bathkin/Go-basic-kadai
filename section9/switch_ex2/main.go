package main

import "fmt"

func main() {
	days := "Sunday"

	switch days {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("平日です")
	case "Saturday", "Sunday":
		fmt.Printf("%vは休日です\n", days)
	default:
		fmt.Println("該当の曜日なし")
	}
}
