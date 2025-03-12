package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	then := time.Date(
		2003, 2, 13, 2, 2, 4, 5, time.UTC)
	fmt.Println(then)

	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())
	fmt.Println(then.Weekday()) // Weekday() is used to get the day of the week.
	fmt.Println(then.Weekday().String())

	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))
	fmt.Println(then.Equal(now)) // =

	diff := now.Sub(then)
	fmt.Println(diff)
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())

	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))
}
