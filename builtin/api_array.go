package main

import "fmt"

func main() {
	// 创建数组，指定长度
	// 类型为[5]int
	var arr1 [5]int
	arr1[0] = 1
	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// key value方式，指定索初始化
	var arr3 = [5]string{3: "third", 4: "last"}
	fmt.Println(arr3)

	arr4 := [...]int{11, 12, 13, 14, 15}
	fmt.Println(arr4)

	arr5 := [...]string{3: "aaa", 4: "xxx"}
	fmt.Println(arr5)

	// 创建数组，new
	// 类型为*[5]int
	var arr6 = new([5]int)
	arr6[0] = 3
	fmt.Println(arr6)
}
