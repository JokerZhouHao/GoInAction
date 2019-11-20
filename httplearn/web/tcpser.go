package main

import (
	"log"
	"net"
)

func dealErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	addr, err := net.ResolveTCPAddr("tcp4", ":8888")
	dealErr(err)
	listener, err := net.ListenTCP("tcp4", addr)
	dealErr(err)
	for {
		conn, err := listener.Accept()
		dealErr(err)
		conn.Write([]byte("周豪"))
		conn.Close()
	}
}
