package main

import (
	"fmt"
	m2 "iotClient/protocol/modbus"
	"testing"
	"time"
)

// 读寄存器
func Test_ReadHoldingRegisters(t *testing.T) {
	var mv m2.TcpClient
	mv.TimeOut = 5 * time.Second
	mv.Address = "localhost:502"

	//init
	err := mv.InitModbus()
	if err != nil {
		panic(err)
	}

	//读寄存器
	values, _ := mv.ReadHoldingRegisters(uint16(99), uint16(3))
	fmt.Println(values)

	mv.Close()
}

// 读线圈
func Test_ReadCoils(t *testing.T) {
	var mv m2.TcpClient
	mv.TimeOut = 5 * time.Second
	mv.Address = "localhost:502"

	//init
	err := mv.InitModbus()
	if err != nil {
		panic(err)
	}

	values, _ := mv.ReadCoils(99, 4)
	fmt.Println(values)
}
