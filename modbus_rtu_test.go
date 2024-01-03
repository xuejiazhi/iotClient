package main

import (
	"fmt"
	m2 "iotClient/protocol/modbus"
	"testing"
)

func initModbusRtuClient() m2.ModbusClient {
	var tm m2.ModbusClient = &m2.RtuClient{
		Address:  "COM2",
		BaudRate: 9600,
		DataBits: 8,
		Parity:   "O",
	}
	//init
	err := tm.InitModbus()
	if err != nil {
		panic(err)
	}
	return tm
}

func Test_ReadRtuHoldingRegisters(t *testing.T) {
	tm := initModbusRtuClient()
	defer tm.Close()
	//读寄存器
	values, _ := tm.ReadHoldingRegisters(uint16(99), uint16(3))
	fmt.Println(values)
}
