package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"

	parseTime, err := time.Parse(layout, "2020-10-24 15:30:30")
	if err != nil {
		panic(err)
	}
	fmt.Println("time: ", parseTime)

	formatTime := time.Now().Format(layout)
	fmt.Println("string: ", formatTime)

}
