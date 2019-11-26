package main

import "fmt"

const (
	PerMinute = 60
	PerHour   = PerMinute * 60
	PerDay    = PerHour * 24
)

func parseTime(seconds int) (day int, hour int, minute int) {
	day = seconds / PerDay
	hour = seconds / PerHour
	minute = seconds / PerMinute

	return
}

func main() {
	fmt.Println(parseTime(300))
	_, hour, _ := parseTime(10000)
	fmt.Println(hour)
}
