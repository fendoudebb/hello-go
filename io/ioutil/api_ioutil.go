package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadAll(strings.NewReader("abcdefg"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
