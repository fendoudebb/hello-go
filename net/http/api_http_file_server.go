package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("get wd error#", err.Error())
	}
	http.Handle("/net/", http.StripPrefix("/net/", http.FileServer(http.Dir(wd + "/net"))))

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(wd + "/builtin"))))

	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal("Listen Fail#", err.Error())
	}
}
