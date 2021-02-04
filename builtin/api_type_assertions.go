package main

import "fmt"

// 类型推断
func main() {
	//var i interface{} = 123
	var i interface{} = "hello"

	var i2 interface{} = int32(1)

	Cast(i)
	Cast(i2)

	//s := i.(string)
	//fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic 即使是字符串123也会panic 因为类型不同
	fmt.Println(f)

}

func Cast(object interface{}) {
	switch object.(type) {
	case string:
		fmt.Println("string#", object.(string))
	case int:
		fmt.Println("int#", object.(int))
	case float64:
		fmt.Println("float64#", object.(float64))
		// 其他类型
	}
}
