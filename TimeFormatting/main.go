package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	fmt.Println(t1)

	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("2006-01-02T15:04:05-07:00"))

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	fmt.Println(t2, e)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	fmt.Println(e)
}
