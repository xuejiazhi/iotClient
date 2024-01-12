package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	fmt.Println("Successfully connected to UDP server")

	//message := "Hello, server!"  BAC0

	addr, err := net.ResolveUDPAddr("udp", "192.168.31.171:47808")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	VirtualLink := []byte{0x81, 0x0b, 0x00, 0x0c}
	NPDU := []byte{0x01, 0x20, 0xff, 0xff, 0x00, 0xff}
	APDU := []byte{0x10, 0x08}
	mearge := append(VirtualLink, append(NPDU, APDU...)...)
	if _, err := conn.Write(mearge); err != nil {
		log.Fatalf("Failed to write Unconfimed request WhoIs packet: %s", err)
	}

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) // 从whois服务器获取返回结果
		if n == 0 || err != nil {
			break
		}
		fmt.Print(string(buf[:n]))
	}

	os.Exit(0)
}
