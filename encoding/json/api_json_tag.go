package main

import (
	"encoding/json"
	"fmt"
)

type (
	R1 struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Data interface{} `json:"data"`
	}

	R2 struct {
		Code int `json:"code,string"`
		Code2 int `json:",string"`
		Msg string `json:"msg,omitempty"`
		Data interface{} `json:",omitempty"`
		Data2 interface{} `json:",omitempty"`
		Id int `json:"-"`
		Id2 int `json:"-,"`
	}
)

func main() {
	// omitempty 空值/默认值 不返回
	// - 忽略字段
	// -,
	// ,string 强转

	r1, _ := json.Marshal(R1{Code: 0, Msg: "请求成功", Data: "字符串数据"})
	fmt.Println(string(r1)) // {"code":0,"msg":"请求成功","data":"字符串数据"}

	r2, _ := json.Marshal(R2{Code: 0, Msg: "", Data: "字符串数据", Data2: nil, Code2: 999, Id: 1000, Id2:2000})
	fmt.Println(string(r2)) // {"code":"0","Code2":"999","Data":"字符串数据","-":2000}

}
