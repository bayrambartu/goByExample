package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// we can set (Setenv), read (Getenv) and list (Environ) environment variables with the os package.
	os.Setenv("FOO", "1") // Create an environment variable named “FOO” and set its value to “1”
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	/* output
	FOO: 1
	BAR:
	*/

	for _, e := range os.Environ() {
		// os.Environ() -> Returns ALL environment variables in KEY=value format.
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
		fmt.Println(pair[1])

	}
	apiKey := os.Getenv("API_KEY")
	fmt.Println("using api key:", apiKey)

	homeDir := os.Getenv("HOME")
	fmt.Println("Home Directory:", homeDir)

	dbPassword := os.Getenv("DB_PASSWORD")
	fmt.Println("DB Password:", dbPassword)

}
