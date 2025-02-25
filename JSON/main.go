package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
	type Person struct {
		Name string
		Age  int
		City string
	}
*/
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `"json:page"`
	Fruits []string `"json:fruits"`
}

func main() {
	/*
		// Convert a Go data structure to JSON format

		person := Person{"bartu", 21, "New York"}
		jsonData, err := json.Marshal(person)
		fmt.Println(jsonData, err) // byte
		if err != nil {
			panic(err)
		}
		fmt.Println("down")
		fmt.Println(string(jsonData))

		// Convert Json format structure to a Go data
		var p2 Person
		jsonStr := []byte(`{"Name": "bartu", "Age": 21, "City": "New York"}`)
		fmt.Println(jsonStr)
		err = json.Unmarshal(jsonStr, &p2)
		if err != nil {
			panic(err)
		}
		fmt.Println(p2)

		fmt.Println("----------")
		jsonString := `{"Name": "bayram", "Age": 31}`
		var data map[string]interface{}
		json.Unmarshal([]byte(jsonString), &data)

		// Demonstrating that JSON is in string format
		fmt.Println("JSON as string:", jsonString)
		fmt.Println("Go map format:", data)

	*/
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	fmt.Println(bolB) // [116 114 117 101]
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	dec := json.NewDecoder(strings.NewReader(`{"page":1, "fruits": ["apple", "peach"]}`))
	res1 := response1{}
	dec.Decode(&res1)
	fmt.Println(res1)
}
