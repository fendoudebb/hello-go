package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "张三")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])

	encode := v.Encode()
	fmt.Println(encode)

	response, err := http.Post("http://www.baidu.com", "application/json", strings.NewReader("aaaa"))

	if err != nil && response !=nil {
		println(response.StatusCode)
	}

}
