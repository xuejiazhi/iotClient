package modbus

import (
	"errors"
	"github.com/goburrow/modbus"
	"iotClient/protocol/comm"
)

/*
*@author
* InitModbus Tcp
 */

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
	if c := GetOperate["checkLessLen"](results, 2); c != nil {
		err = c.(error)
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
// address 地址
// value 值
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

// WriteMultipleRegisters 批量写入寄存器
// address 地址
// quantity 数量
// values 更新的值列表
func (t *TcpClient) WriteMultipleRegisters(address, quantity uint16, values []int) (err error) {
	//check eq len
	if c := GetOperate["checkEqLen"](len(values), int(quantity)); c != nil {
		err = c.(error)
		return
	}
	//将INT换算成Bytes
	dataBytes := intToBytes(values)
	if c := GetOperate["checkEqLen"](len(dataBytes), 2*int(quantity)); c != nil {
		err = c.(error)
		return
	}
	//写入data
	_, err = t.Client.WriteMultipleRegisters(address, quantity, dataBytes)

	//return
	return
}

// WriteSingleCoil 写入单个线圈
func (t *TcpClient) WriteSingleCoil(address, value uint16) (err error) {
	//是否在数组里面
	if !comm.InIntArray(int(value), []int{0, 1}) {
		err = errors.New("modbus: state '1' must be either 1 (ON) or 0 (OFF)")
		return
	}

	//取值
	coilValue := comm.If(value == 1, CoilStateOn, CoilStateOff).(int)
	//写入值
	_, err = t.Client.WriteSingleCoil(address, uint16(coilValue))
	if err != nil {
		return
	}

	//return data
	return
}

func (t *TcpClient) WriteMultipleCoils(address, quantity uint16, value []byte) {

}
