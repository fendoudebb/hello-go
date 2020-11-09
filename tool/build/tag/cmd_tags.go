package main

import "fmt"

// go build -tags dev
// 执行带tags属性的build命令时，必须注释了，否则会报重复定义
var version = "dev"

func main() {
	fmt.Printf("running %s version", version)
}
