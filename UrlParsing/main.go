package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)     // '/path'
	fmt.Println(u.Fragment) //section of  start with '#' -> f

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

	fmt.Println("---------------------------")
	a := "https://example.com/search?q=golang&lang=en"
	fmt.Println(a)
	b, _ := url.Parse(a)
	fmt.Println(b)
	fmt.Println(b.Host)
	hosst, porrt, _ := net.SplitHostPort(a)
	fmt.Println(hosst)
	fmt.Println(porrt)
	c, _ := url.ParseQuery(b.RawQuery)
	fmt.Println(c)
	fmt.Println(c["q"][0])
	fmt.Println(c["lang"][0])

}
