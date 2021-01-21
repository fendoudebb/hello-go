package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	defer server.Close()
	if err != nil {
		log.Fatal("err#" + err.Error())
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("conn error#" + err.Error())
			continue
		}

		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("remote addr#" + conn.RemoteAddr().String())

	buf := make([]byte, 2)

	for  {
		fmt.Println("read...")
		length, err := conn.Read(buf)
		if err != nil || length == 0{
			log.Println("err != nil or length = 0")
			break
		}

		bytes := buf[:length]
		s := string(bytes)
		fmt.Println(s)

		conn.Write([]byte("nihao"))
	}

}
