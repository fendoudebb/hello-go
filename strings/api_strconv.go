package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	atoi, err := strconv.Atoi("20201025")
	fmt.Println(atoi, err)

	itoa := strconv.Itoa(20201024)
	fmt.Println(itoa)

	// 32位float，64位float
	float, err := strconv.ParseFloat("2020.1025", 64)
	fmt.Println(float, err)

	var age = 18
	fmt.Println(strconv.Itoa(age))

	unixTimestamp := time.Now().Unix()
	fmt.Println(strconv.FormatInt(unixTimestamp, 10))


}
