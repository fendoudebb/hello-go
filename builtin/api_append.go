package main

import (
	"fmt"
	"reflect"
)

func main() {

	var arr  = []int{1, 2, 3}
	var arr2  = [3]int{1, 2, 3}

	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(arr2))

	ints := append(arr, 1)
	ints[0] = 100
	fmt.Println(arr)
	fmt.Println(ints)

}
