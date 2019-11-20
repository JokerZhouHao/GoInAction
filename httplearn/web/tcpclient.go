package main

import (
	"fmt"
	"log"
	"net"
)

func dealErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", ":8888")
	//dealErr(err)
	conn, _ := net.DialTCP("tcp4", nil, addr)
	//dealErr(err)
	bs := make([]byte, 1024)
	num, _ := conn.Read(bs)
	//dealErr(err)
	fmt.Println(string(bs[:num]))
	conn.Close()
}
