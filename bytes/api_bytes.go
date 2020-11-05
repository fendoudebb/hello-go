package main

import (
	"bytes"
	"fmt"
)

func main() {
	bufferString := bytes.NewBufferString("abc")
	bufferString.WriteString("efg")
	fmt.Println(bufferString.String())

}
