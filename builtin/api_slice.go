package main

import "fmt"

func main() {

	// 根据数组获取切片
	arr := [...]int{1, 2, 3,4,5}
	// 获取与数组等长的切片
	s := arr[:]
	// [1 2 3 4 5]
	fmt.Println(s)

	// 等同于前一个方式，直接获取与数组等长的切片
	slice1 := []int{6, 7, 8, 9, 10}
	// [6 7 8 9 10]
	fmt.Println(slice1)

	// 按索引初始化
	slice2 := []string{5: "eee", 2: "bbb"}
	// [  bbb   eee]
	fmt.Println(slice2)
	// 6
	fmt.Println(len(slice2))
	// 6
	fmt.Println(cap(slice2))

	// 长度为5，容量为10
	ints := make([]int, 5, 10)
	fmt.Println(ints)
	// 5
	fmt.Println(len(ints))
	// 10
	fmt.Println(cap(ints))

	var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
	var a = ar[10:]
	// 0
	fmt.Println(len(a))
	// 0
	fmt.Println(cap(a))

}
