package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	/*
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return

		}
		fmt.Println("data from the server", string(body)) */

	/*
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		req.Header.Set("User-Agent", "MyCustomClient/1.0") // add header

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("response not read", err)
			return
		}
		fmt.Println("incoming data:", string(body))
	*/

	/*
		jsonData := `{"title": "Golang HTTP Client", "body": "HTTP is fun!", "userId": 1}`
		resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer([]byte(jsonData)))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("answer not read", err)
			return
		}
		fmt.Println("response from server ", string(body))

	*/

	/*
		// Send POST Request with http.NewRequest
		client := &http.Client{}
		jsonData := `{"title": "Golang", "body": "Advanced HTTP Client", "userId": 1}`
		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer([]byte(jsonData)))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer your_token_here")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("answer not read", err)
			return
		}
		fmt.Println("response from server:", string(body))
	*/
	/*
		client := &http.Client{}
		jsonData := `{"title": "Golang", "body": "Advanced HTTP Client", "userId": 1}`
		req, _ := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer([]byte(jsonData)))
		req.Header.Set("Content-Type", "application/json")
		client.Do(req)
	*/
	/*
		client := &http.Client{}
		req, _ := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/1", nil)
		client.Do(req)
	*/

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)

	}

}
