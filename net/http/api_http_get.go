package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.baidu.com")
	defer func() {
		if res != nil {
			err = res.Body.Close()
			if err != nil {
				log.Fatal("close error#", err)
			}
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	str, _ := ioutil.ReadAll(res.Body)
	log.Println(string(str))

}
