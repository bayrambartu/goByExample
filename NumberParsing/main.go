package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconv
	f, _ := strconv.ParseFloat("44", 32)
	fmt.Printf("type : %T, value: :%v", f, f)
	fmt.Println()
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	j, _ := strconv.ParseInt("1101", 2, 64)
	fmt.Println(j)
	k, _ := strconv.ParseInt("1A", 16, 64)
	fmt.Println(k)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	v, _ := strconv.Atoi("7")
	fmt.Println(v)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
	c, _ := strconv.Atoi("wat")
	fmt.Println(c)

	value, err := strconv.Atoi("123abc")
	if err != nil {
		fmt.Println("error:", err)

	} else {
		fmt.Println("value:", value)
	}

}
