package main

import (
	"fmt"
	"github.com/goburrow/modbus"
	"log"
	"time"
)

type ModBusClient struct {
	Client     modbus.Client
	TimeOut    time.Duration
	DeviceId   byte         //设备ID
	TcpInfo    ModBusTcp    //TCP
	SerialInfo ModBusSerial //串口
}

type ModBusTcp struct {
	Address string //TCP 地址 localhost:502
	Port    int    //TCP 端口
}

// ModBusSerial 串口配置
type ModBusSerial struct {
	CommAddr string //串口地址 COM1,COM2
	BaudRate int    //波特率
	DeviceID int    //设备ID
	DataBits int
	Parity   string
	StopBits int
}

// InitTcpModbus Tcp Modbus
func (m ModBusClient) InitTcpModbus() (err error) {
	//set DeviceID
	if m.DeviceId == 0 {
		m.DeviceId = 1
	}
	//Address And Port
	if m.TcpInfo.Address == "" {
		m.TcpInfo.Address = "localhost"
	}
	if m.TcpInfo.Port == 0 {
		m.TcpInfo.Port = 502
	}

	//set TcpAddr
	tcpAddr := fmt.Sprintf("%s%d", m.TcpInfo.Address, m.TcpInfo.Port)
	//set TcpHandle
	tcpHandler := modbus.NewTCPClientHandler(tcpAddr)
	//connect
	if err = tcpHandler.Connect(); err != nil {
		return
	}
	//set deviceId
	tcpHandler.SlaveId = m.DeviceId
	//set timeout
	tcpHandler.Timeout = m.TimeOut
	//set Modbus Client
	m.Client = modbus.NewClient(tcpHandler)
	//return
	return
}

func (m ModBusClient) InitSerialModbus() (err error) {
	//set DeviceID
	if m.DeviceId == 0 {
		m.DeviceId = 1
	}
	//设置串口
	if m.SerialInfo.CommAddr == "" {
		m.SerialInfo.CommAddr = "COM1"
	}
	//设置波特率
	if m.SerialInfo.BaudRate == 0 {
		m.SerialInfo.BaudRate = 9600
	}
	if m.SerialInfo.DataBits == 0 {
		m.SerialInfo.DataBits = 8
	}
	if m.SerialInfo.StopBits == 0 {
		m.SerialInfo.StopBits = 1
	}
	if m.SerialInfo.Parity == "" {
		m.SerialInfo.Parity = "O"
	}

	//set Serial Address
	serialAddr := m.SerialInfo.CommAddr
	serialHandler := modbus.NewRTUClientHandler(serialAddr)

	//connect
	if err = serialHandler.Connect(); err != nil {
		return
	}

	serialHandler.SlaveId = m.DeviceId
	serialHandler.BaudRate = m.SerialInfo.BaudRate
	serialHandler.DataBits = m.SerialInfo.DataBits
	serialHandler.Parity = m.SerialInfo.Parity
	serialHandler.StopBits = m.SerialInfo.StopBits
	if m.TimeOut.Seconds() > 0 {
		serialHandler.Timeout = m.TimeOut
	}

	//set Client
	m.Client = modbus.NewClient(serialHandler)
	return
}

func InitModbus() {
	tcpHandler := modbus.NewTCPClientHandler("127.0.0.1:502")
	err := tcpHandler.Connect()
	tcpHandler.SlaveId = 2
	if err != nil {
		log.Fatal(err)
	}
	defer tcpHandler.Close()
	client := modbus.NewClient(tcpHandler)
	//results, err := client.ReadHoldingRegisters(99, 66)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////
	//a := []byte{0x00, 0x48, 0x00, 0x66}
	//client.WriteMultipleRegisters(99, 2, a)
	//log.Printf("%v\n", results)

	resultR, err := client.ReadCoils(99, 3)
	fmt.Println(resultR)
}

// InitModbusComm 串口
func InitModbusComm() {
	addr := "COM2"
	handler := modbus.NewRTUClientHandler(addr)
	handler.SlaveId = 1
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "O"
	handler.StopBits = 1
	handler.Timeout = 5 * time.Second
	defer handler.Close()

	client := modbus.NewClient(handler)
	//address := uint16(0x0060)
	//quantity := uint16(0x0003)
	results, err := client.ReadHoldingRegisters(99, 2)
	fmt.Println(results, err)
	//ClientTestAll(t, modbus.NewClient(handler))
}
