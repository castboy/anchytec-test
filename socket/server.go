package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	fmt.Println(err)

	conn, err := listener.Accept()
	fmt.Println(err)

	conn.SetReadDeadline(time.Now().Add(time.Duration(100) * time.Second))
	data := make([]byte, 100)
	for {
		n, err := conn.Read(data)
		fmt.Println(n, err, string(data))
	}
}
