package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("pinch"))
	fmt.Println(r.MatchString("pe1ach"))

	fmt.Println(r.FindString("pinch punch peach"))
	fmt.Println(r.FindStringIndex("pjsch"))
	fmt.Println(r.FindStringSubmatch(" pinch peach"))
	fmt.Println(r.FindAllString("pinch peach", -1))

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// EmailRegex
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	fmt.Println(emailRegex.MatchString("test@example.com"))
	fmt.Println(emailRegex.MatchString("kurnazb32.com"))

	// Phone Regex
	phoneRegex := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	fmt.Println(phoneRegex.MatchString("+1234567890"))
	fmt.Println(phoneRegex.MatchString("123-456-789"))

	// web
	urlRegex := regexp.MustCompile(`https?://[^\s]+`)
	text := "Check out https://golang.org and http://example.com"
	fmt.Println(urlRegex.FindAllString(text, -1))
}
