package modbus

import (
	"github.com/goburrow/modbus"
	"iotClient/protocol/comm"
	"time"
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

	WriteSingleRegister(address, value uint16) error
}

/*
*@author
* InitModbus Tcp
 */

type TcpClient struct {
	SlaveId byte
	Address string //TCP 地址 localhost:502
	TimeOut time.Duration
	Client  modbus.Client
	Handler *modbus.TCPClientHandler
}

func (t *TcpClient) InitModbus() (err error) {
	//Address And Port
	t.Address = GetOperate["tcpAddr"](t.Address, "localhost:502").(string)
	t.Handler = modbus.NewTCPClientHandler(t.Address)

	//connect
	if err = t.Handler.Connect(); err != nil {
		return
	}

	//set DeviceID
	t.Handler.SlaveId = GetOperate["slaveId"](t.SlaveId, uint8(1)).(byte)
	// set timeout
	if t.TimeOut.Seconds() > 0 {
		t.Handler.Timeout = t.TimeOut
	}
	//set client
	t.Client = modbus.NewClient(t.Handler)

	//return
	return
}

func (t *TcpClient) Close() (err error) {
	return t.Handler.Close()
}

// ReadHoldingRegisters 读取寄存器数据
// address 地址
// quantity 数量
func (t *TcpClient) ReadHoldingRegisters(address, quantity uint16) (values []int, err error) {
	//读取寄存器
	results, err := t.Client.ReadHoldingRegisters(address, quantity)
	//check error
	if err != nil {
		return
	}
	//check less len
	if err = GetOperate["checkLessLen"](results, 2).(error); err != nil {
		return
	}

	//设置数据
	for i := 0; i < len(results); i = i + 2 {
		//一个数据为两个byte
		dataBytes := results[i : i+2]
		if len(dataBytes) == 2 {
			values = append(values, getRegisterValue(dataBytes))
		}
	}

	//return data
	return
}

// ReadCoils 读取线圈
// address 地址
// quantity 数量
func (t *TcpClient) ReadCoils(address, quantity uint16) (data []int, err error) {
	// 读点位线圈数据
	results, err := t.Client.ReadCoils(address, quantity)
	if err != nil {
		return
	}
	//check less len
	if err = GetOperate["checkLessLen"](results, 1).(error); err != nil {
		return
	}

	//取data
	data = comm.DecimalToBinary(int(results[0]))

	//return data
	return
}

// ReadInputStatus 输入状态
// address 地址
// quantity 数量
func (t *TcpClient) ReadInputStatus(address, quantity uint16) (data []int, err error) {
	// 读点位线圈数据
	results, err := t.Client.ReadDiscreteInputs(address, quantity)
	//check error
	if err != nil {
		return
	}
	//check less len
	if err = GetOperate["checkLessLen"](results, 1).(error); err != nil {
		return
	}

	//取data
	data = comm.DecimalToBinary(int(results[0]))

	//return data
	return
}

// ReadInputRegisters 输入寄存器
// address 地址
// quantity 数量
func (t *TcpClient) ReadInputRegisters(address, quantity uint16) (values []int, err error) {
	//读取寄存器
	results, err := t.Client.ReadInputRegisters(address, quantity)
	if err != nil {
		return
	}

	//设置数据
	for i := 0; i < len(results); i = i + 2 {
		//一个数据为两个byte
		dataBytes := results[i : i+2]
		if len(dataBytes) == 2 {
			values = append(values, int(dataBytes[0])*256+int(dataBytes[1]))
		}
	}

	//return data
	return
}

// WriteSingleRegister 写入单个寄存器
func (t *TcpClient) WriteSingleRegister(address, value uint16) (err error) {
	//写入单个寄存器
	result, err := t.Client.WriteSingleRegister(address, value)
	if err != nil {
		return
	}
	//check less len
	if err = GetOperate["checkLessLen"](result, 2).(error); err != nil {
		return
	}
	//return data
	return
}
