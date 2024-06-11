package main

import (
	"fmt"
	"time"
)

func init() {
	// Load Asia/Jakarta timezone
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	time.Local = loc
}

func main() {
	fmt.Println("Current time in Jakarta:", time.Now().Format("2006-01-02 15:04"))
}
