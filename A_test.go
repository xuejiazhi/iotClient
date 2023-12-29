package main

import (
	"fmt"
	"github.com/goburrow/serial"
	"testing"
)

func TestInitModbus(t *testing.T) {
	cfg := &serial.Config{Address: "COM1", BaudRate: 9600, Timeout: 3 /*毫秒*/}

	_, err := serial.Open(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
}
