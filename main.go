package main

import (
	"fmt"
	"time"
)

// functtion to return yesterday datetime
func yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

func main() {
	fmt.Println("Yesterday's date is:", yesterday().Format("2006-01-02"))
	fmt.Println("Hello, World!")
}
