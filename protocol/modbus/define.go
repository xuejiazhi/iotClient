package modbus

import (
	"github.com/goburrow/modbus"
	"time"
)

var (
	CoilStateOn  = 0xFF00
	CoilStateOff = 0x0000
)

/**
01：COIL STATUS（线圈状态）：用于读取和控制远程设备的开关状态，通常用于控制继电器等开关设备。
02：INPUT STATUS（输入状态）：用于读取远程设备的输入状态，通常用于读取传感器等输入设备的状态。
03：HOLDING REGISTER（保持寄存器）：用于存储和读取远程设备的数据，通常用于存储控制参数、设备状态等信息。
04：INPUT REGISTER（输入寄存器）：用于存储远程设备的输入数据，通常用于存储传感器等输入设备的数据。
*/

type ModbusClient interface {
	InitModbus() error
	Close() error
	ReadHoldingRegisters(uint16, uint16) ([]int, error)
	ReadCoils(uint16, uint16) ([]int, error)
	ReadInputStatus(uint16, uint16) ([]int, error)
	ReadInputRegisters(uint16, uint16) ([]int, error)

	WriteSingleRegister(uint16, uint16) error
	WriteMultipleRegisters(address, quantity uint16, values []int) error
	WriteSingleCoil(uint16, uint16) error
	WriteMultipleCoils(uint16, uint16, []int) error
}

// TcpClient Tcp
type TcpClient struct {
	SlaveId byte
	Address string //TCP 地址 localhost:502
	TimeOut time.Duration
	Client  modbus.Client
	Handler *modbus.TCPClientHandler
}

// RtuClient 串口
type RtuClient struct {
	SlaveId  byte
	Address  string
	BaudRate int64  //波特率
	DataBits int8   //数据位
	Parity   string //奇偶校验位 O:奇; E:偶
	StopBits int
	TimeOut  time.Duration
	Client   modbus.Client
	Handler  *modbus.RTUClientHandler
}
