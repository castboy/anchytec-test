package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":7777")
	fmt.Println(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	fmt.Println(err)
	for {
		_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
		fmt.Println(err)
		time.Sleep(5 * time.Second)
	}
}
