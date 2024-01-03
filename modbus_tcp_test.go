package main

import (
	"fmt"
	m2 "iotClient/protocol/modbus"
	"log"
	"testing"
	"time"
)

func initModbusClient() m2.ModbusClient {
	var tm m2.ModbusClient = &m2.TcpClient{
		TimeOut: 5 * time.Second,
		Address: "localhost:502",
	}
	//init
	err := tm.InitModbus()
	if err != nil {
		panic(err)
	}
	return tm
}

// 读寄存器
func Test_ReadHoldingRegisters(t *testing.T) {
	tm := initModbusClient()
	defer tm.Close()

	//读寄存器
	values, _ := tm.ReadHoldingRegisters(uint16(99), uint16(3))
	fmt.Println(values)

}

// 读线圈
func Test_ReadCoils(t *testing.T) {
	tm := initModbusClient()
	defer tm.Close()

	//读线圈
	values, _ := tm.ReadCoils(99, 4)
	fmt.Println(values)
}

func Test_ReadInputStatus(t *testing.T) {
	tm := initModbusClient()
	defer tm.Close()

	//读输入状态
	values, _ := tm.ReadInputStatus(99, 4)
	fmt.Println(values)
}

// 读输入寄存器
func Test_ReadInputRegister(t *testing.T) {
	tm := initModbusClient()
	defer tm.Close()

	//读输入寄存器
	values, _ := tm.ReadInputRegisters(uint16(99), uint16(3))
	fmt.Println(values)

	tm.Close()
}

// 写入寄存器
func Test_WriteSingleRegister(t *testing.T) {
	tm := initModbusClient()
	defer tm.Close()

	if err := tm.WriteSingleRegister(99, 120); err != nil {
		log.Fatal(err)
	}
}

// 批量写入寄存器
func Test_WriteMultipleRegisters(t *testing.T) {
	tm := initModbusClient()
	tm.WriteMultipleRegisters(99, 4, []int{4, 3, 2, 1})
	tm.Close()
}

// 写入线圈
func Test_WriteCoils(t *testing.T) {
	tm := initModbusClient()
	tm.WriteSingleCoil(99, 1)
	tm.Close()
}

// 批量写线圈
func Test_WriteMultiCoils(t *testing.T) {
	tm := initModbusClient()
	tm.WriteMultipleCoils(99, 5, []int{1, 1, 1, 0, 0})
	tm.Close()
}
