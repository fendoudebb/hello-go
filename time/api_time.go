package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("time.Now()：", t)
	year, month, day := t.Date()
	fmt.Println("日期：", year, month, day) // 2020 October 24
	fmt.Println("一年中的第几天：", t.YearDay())
	fmt.Println("星期几：", t.Weekday()) // Saturday
	fmt.Println("年：", t.Year())
	fmt.Println("月：", t.Month()) // October
	fmt.Println("日：", t.Day())
	fmt.Println("时：", t.Hour())
	fmt.Println("分：", t.Minute())
	fmt.Println("秒：", t.Second())
	fmt.Println("纳秒：", t.Nanosecond())
	fmt.Println("秒时间戳：", t.Unix())
	fmt.Println("纳秒时间戳：", t.UnixNano())
	fmt.Println("毫秒时间戳：", t.UnixNano() / 1e6)

	secondToTime := time.Unix(1603546715, 0)
	fmt.Println("秒值转Time：", secondToTime) // 2020-10-24 21:38:35 +0800 CST

	nanoSecondToTime := time.Unix(0, 1603546715761482000)
	fmt.Println("毫秒值转Time：", nanoSecondToTime) // 2020-10-24 21:38:35.761482 +0800 CST

	fmt.Println("格式化：", t.Format(time.Kitchen)) // 9:42PM

	time.Sleep(2 * time.Second) // 休眠2秒
	time.Sleep(2e9) // 休眠2秒

	delta := time.Now().Sub(t)
	fmt.Println("时间差：", delta) // 4.0534341s

}
