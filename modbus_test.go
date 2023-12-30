package main

import (
	"fmt"
	m2 "iotClient/protocol/modbus"
	"testing"
	"time"
)

// 读寄存器
func Test_ReadHoldingRegisters(t *testing.T) {
	var tm m2.ModbusClient = &m2.TcpClient{
		TimeOut: 5 * time.Second,
		Address: "localhost:502",
	}
	//init
	err := tm.InitModbus()
	if err != nil {
		panic(err)
	}

	//读寄存器
	values, _ := tm.ReadHoldingRegisters(uint16(99), uint16(3))
	fmt.Println(values)

	tm.Close()
}

// 读线圈
func Test_ReadCoils(t *testing.T) {
	var tm m2.ModbusClient = &m2.TcpClient{
		TimeOut: 5 * time.Second,
		Address: "localhost:502",
	}

	err := tm.InitModbus()
	defer tm.Close()
	if err != nil {
		panic(err)
	}

	//读线圈
	values, _ := tm.ReadCoils(99, 4)
	fmt.Println(values)
}

func Test_ReadInputStatus(t *testing.T) {
	var tm m2.ModbusClient = &m2.TcpClient{
		TimeOut: 5 * time.Second,
		Address: "localhost:502",
	}

	err := tm.InitModbus()
	defer tm.Close()
	if err != nil {
		panic(err)
	}

	//读输入状态
	values, _ := tm.ReadInputStatus(99, 4)
	fmt.Println(values)
}

// 读输入寄存器
func Test_ReadInputRegister(t *testing.T) {
	var tm m2.ModbusClient = &m2.TcpClient{
		TimeOut: 5 * time.Second,
		Address: "localhost:502",
	}
	//init
	err := tm.InitModbus()
	if err != nil {
		panic(err)
	}

	//读输入寄存器
	values, _ := tm.ReadInputRegisters(uint16(99), uint16(3))
	fmt.Println(values)

	tm.Close()
}
