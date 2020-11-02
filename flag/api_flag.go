package main

import (
	"flag"
	"fmt"
)

var s string

var show bool

func init() {
	// go run api_flag.go -test_filed="hello world"
	// go run api_flag.go -test_filed "hello world"
	flag.StringVar(&s, "test_filed", "this is default value", "this is description")

	// go run api_flag.go -show
	// go run api_flag.go -show=true
	// go run api_flag.go -show=t
	// go run api_flag.go -show=T
	// go run api_flag.go -show=false
	// go run api_flag.go -show=f
	// go run api_flag.go -show=F
	// go run api_flag.go -show=0
	// go run api_flag.go -show=1
	flag.BoolVar(&show, "show", false, "")

/*
	-flag // 只支持bool类型
	-flag=x
	-flag x // 只支持非bool类型
*/

}

func main() {
	flag.Parse()

	fmt.Println(s)

	fmt.Println(show)

	fmt.Println("预设的flag的参数个数：", flag.NFlag())

	// go run api_flag.go hello world test1 test2
	fmt.Println("非预设的flag的参数：", flag.Args()) // [hello world test1 test2]
	fmt.Println(flag.Arg(0)) // hello
	fmt.Println("非预设的flag的参数个数：", flag.NArg())

}
