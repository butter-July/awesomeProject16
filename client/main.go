package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:8888")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	reader := bufio.NewReader(os.Stdin)
	for {
		information, _, _ := reader.ReadLine()
		conn.Write(information)
		br := make([]byte, 1024)
		rn, _ := conn.Read(br)
		fmt.Println(string(br[:rn]))
	}
}
