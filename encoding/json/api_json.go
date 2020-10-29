package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func main() {

	addressStruct := &Address{`CN`, "Shanghai", "China"}

	marshal, err := json.Marshal(addressStruct)
	if err != nil {
		log.Fatal(err)
	}
	// {"Type":"CN","City":"Shanghai","Country":"China"}
	fmt.Printf("Address JSON format: %s\n", marshal)


	result := &Result{0, "请求成功", "this is data"}
	resultStr, _ := json.Marshal(result)
	// {"code":0,"msg":"请求成功","data":"this is data"}
	fmt.Printf("Result JSON format: %s\n", resultStr)

	m := make(map[string]interface{})
	m["id"] = 1
	m["price"] = 10.24
	m["vip"] = true
	m["create_ts"] = "2020-10-28"
	m["remark"] = nil
	mapStr, _ := json.Marshal(m)
	// {"create_ts":"2020-10-28","id":1,"price":10.24,"remark":null,"vip":true}
	fmt.Printf("Map JSON format: %s\n", mapStr)


	addressJsonStr := `{"Type":"CN","City":"Shanghai","Country":"China"}`
	var address Address
	err = json.Unmarshal([]byte(addressJsonStr), &address)
	if err != nil {
		log.Println(err)
	}
	// {CN Shanghai China}
	fmt.Println(address)

	m2 := make(map[string]interface{})
	err = json.Unmarshal([]byte(addressJsonStr), &m2)
	// map[City:Shanghai Country:China Type:CN]
	fmt.Println(m2)

	const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`
	var i []map[string]interface{}

	decoder := json.NewDecoder(strings.NewReader(jsonStream))
	err = decoder.Decode(&i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)

}

type User struct {
	Name string
	Age int
}


// json.NewDecoder(r.Body)
func HandleUse(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "姓名：%s，年龄：%d", u.Name, u.Age)
}
