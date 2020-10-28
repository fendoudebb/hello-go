package main

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
	file  *os.File
)

func init() {
	f, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("open log file err: ", err)
	}
	file = f

	log.SetPrefix("DEBUG: ")
	flag := log.Ldate | log.Lmicroseconds | log.Lshortfile
	log.SetFlags(flag)
	Info = log.New(file, "INFO: ", flag)
	Warn = log.New(file, "WARN: ", flag)
	Error = log.New(file, "ERROR: ", flag)
}

func main() {
	defer func() {
		_ = file.Close()
	}()
	log.Println("this is a service log")
	Info.Println("this is info level log")
	Warn.Println("this is warn level log")
	Error.Println("this is error level log")

}
