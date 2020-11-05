package main

import "fmt"

func main() {

	var m = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	delete(m, "key1")

	// map[key2:value2]
	fmt.Println(m)

}
