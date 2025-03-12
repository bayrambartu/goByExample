package main

import (
	"fmt"
	"time"
)

func main() {
	// Unix Time
	/*
		now := time.Now()
		fmt.Println(now)

		fmt.Println(now.Unix())
		fmt.Println(now.UnixMilli())
		//fmt.Println(now.UnixMicro())
		fmt.Println(now.UnixNano())

		fmt.Println(time.Unix(now.Unix(), 0))
		fmt.Println(time.Unix(0, now.UnixNano()))*/

	now := time.Now()
	fmt.Println("------------")
	fmt.Println(now.Unix())
	tenSecondsLater := now.Add(time.Second * 10)
	fmt.Println("now:", now)
	fmt.Println("10 seconds later:", tenSecondsLater)

	apiTimestamp := int64(1707572730)
	convertedTime := time.Unix(apiTimestamp, 0)
	fmt.Println(apiTimestamp)
	fmt.Println(convertedTime)
	fmt.Println(convertedTime.Format("2006-01-02 15:04:05"))

	start := time.Now()
	time.Sleep(3 * time.Second)
	elapsed := time.Since(start)
	fmt.Println(elapsed.Seconds()) // 3 s
}
