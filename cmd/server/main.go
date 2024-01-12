package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.PacketConn) {
	fmt.Println("New UDP client connected")

	buf := make([]byte, 1024)
	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		fmt.Println("Error reading from UDP client", err.Error())
		return // 终止程序
	}

	fmt.Println("Received message from UDP client:", string(buf[:n]))

	response := "Hello, client!"
	_, err = conn.WriteTo([]byte(response), addr)
	if err != nil {
		fmt.Println("Error writing to UDP client", err.Error())
		return // 终止程序
	}
}

func main() {
	listener, err := net.ListenPacket("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // 终止程序
	}
	defer listener.Close()

	fmt.Println("UDP server listening on port 8080")

	for {
		buf := make([]byte, 1024)
		_, addr, err := listener.ReadFrom(buf)
		if err != nil {
			fmt.Println("Error reading from UDP client", err.Error())
			return // 终止程序
		}

		go handleConnection(listener)

		_, err = listener.WriteTo([]byte("Hello,1111111111111111111111 client!"), addr)
		if err != nil {
			fmt.Println("Error writing to UDP client", err.Error())
			return // 终止程序
		}
	}
}
