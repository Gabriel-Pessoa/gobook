package main

import (
	"fmt"
	"time"
)

func main() {
	const layout = "2006-01-02"
	date := "2022-03-23"

	convertedTime, _ := time.Parse(layout, date)

	fmt.Println(convertedTime)

	fmt.Println(getDaysFromDate(convertedTime))
	fmt.Println(getDaysFromDate(time.Now()))
}

func getDaysFromDate(date time.Time) int {
	diff := time.Since(date)
	days := int(diff.Hours() / 24)
	return days
}
