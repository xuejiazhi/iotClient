package modbus

import (
	"errors"
	"github.com/goburrow/modbus"
	"iotClient/protocol/comm"
	"time"
)

/*
*@author
* InitModbus Tcp
 */

type Operate func(x, y any) any

func GetModBusSlaveID(x, y any) any { return comm.If(x.(uint8) > 0, x, y).(uint8) }
func GetModBusTcpAddr(x, y any) any { return comm.If(x != "", x, y).(string) }
func GetModBusTcpPort(x, y any) any { return comm.If(x.(int) > 0, x, y).(int) }

var GetOperate = map[string]Operate{
	"slaveId": GetModBusSlaveID,
	"tcpAddr": GetModBusTcpAddr,
	"tcpPort": GetModBusTcpPort,
}

type TcpClient struct {
	Client  modbus.Client
	Handler *modbus.TCPClientHandler
	TimeOut time.Duration
	SlaveId byte
	Address string //TCP 地址 localhost:502
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

// ReadCoils 读取线圈
func (t *TcpClient) ReadCoils(address, quantity uint16) (data []int, err error) {
	// 读点位线圈数据
	results, err := t.Client.ReadCoils(address, quantity)
	if err != nil {
		return
	}
	//取data
	if len(results) == 1 {
		data = comm.DecimalToBinary(int(results[0]))
	} else {
		err = errors.New("get Coils data nil")
	}

	//return data
	return
}
