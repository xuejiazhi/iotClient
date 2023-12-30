package main

import (
	"fmt"
	m2 "iotClient/protocol/modbus"
	"testing"
	"time"
)

func Test_ModbusTcp(t *testing.T) {
	var mv m2.TcpClient
	mv.TimeOut = 5 * time.Second
	mv.Address = "localhost:502"

	//init
	err := mv.InitModbus()
	if err != nil {
		panic(err)
	}

	values, _ := mv.ReadHoldingRegisters(uint16(99), uint16(3))
	fmt.Println(values)

	mv.Close()
}
