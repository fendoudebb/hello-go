package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("time.Now()：", t) // 2020-10-24 22:10:53.328973 +0800 CST m=+0.006015101
	year, month, day := t.Date()
	fmt.Println("日期：", year, month, day) // 2020 October 24
	fmt.Println("一年中的第几天：", t.YearDay()) // 298
	fmt.Println("星期几：", t.Weekday()) // Saturday
	fmt.Println("年：", t.Year()) // 2020
	fmt.Println("月：", t.Month()) // October
	fmt.Println("日：", t.Day()) // 24
	fmt.Println("时：", t.Hour()) // 22
	fmt.Println("分：", t.Minute()) // 10
	fmt.Println("秒：", t.Second()) // 53
	fmt.Println("纳秒：", t.Nanosecond()) // 328973000
	fmt.Println("秒时间戳：", t.Unix()) // 1603548653
	fmt.Println("纳秒时间戳：", t.UnixNano()) // 1603548653328973000
	fmt.Println("毫秒时间戳：", t.UnixNano() / 1e6) // 1603548653328

	secondToTime := time.Unix(1603546715, 0)
	fmt.Println("秒值转Time：", secondToTime) // 2020-10-24 21:38:35 +0800 CST

	nanoSecondToTime := time.Unix(0, 1603546715761482000)
	fmt.Println("毫秒值转Time：", nanoSecondToTime) // 2020-10-24 21:38:35.761482 +0800 CST

	fmt.Println("格式化：", t.Format(time.Kitchen)) // 10:10PM

	delta := time.Now().Sub(t)
	fmt.Println("时间差：", delta) // 4.0534341s

	addOneHour := t.Add(time.Hour)
	addTwoHour := t.Add(2 * time.Hour)
	fmt.Println("增加1小时：", addOneHour)
	fmt.Println("增加2小时：", addTwoHour)

	subTwoHour := t.Add(-2 * time.Hour)
	fmt.Println("减去2小时：", subTwoHour)

	addDate := t.AddDate(1, 0, 0)
	fmt.Println("增加1年：", addDate) // 2021-10-24 22:10:53.328973 +0800 CST

	subDate := t.AddDate(-1, 0, 0)
	fmt.Println("减去1年：", subDate) // 2019-10-24 22:10:53.328973 +0800 CST

	before := t.Before(t.Add(time.Hour))
	fmt.Println("before：", before)

	after := t.After(t.Add(time.Hour))
	fmt.Println("after：", after)

	time.Sleep(2 * time.Second) // 休眠2秒
	time.Sleep(2e9) // 休眠2秒

}
