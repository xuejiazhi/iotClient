package modbus

import (
	"errors"
	"github.com/goburrow/modbus"
	"iotClient/protocol/comm"
	"runtime"
)

/**
串口读取
*/

func (r *RtuClient) InitModbus() (err error) {
	//判断OS
	if r.Address == "" {
		if r.Address, err = func() (serial string, err error) {
			os := runtime.GOOS
			switch os {
			case "windows":
				serial = "COM1"
			case "linux":
				serial = "/dev/ttyUSB0"
			default:
				err = errors.New("OS not support")
			}
			return
		}(); err != nil {
			return
		}
	}

	//Address Handler
	r.Handler = modbus.NewRTUClientHandler(r.Address)
	r.Handler.SlaveId = GetOperate["slaveId"](r.SlaveId, uint8(1)).(byte)
	r.Handler.Parity = comm.If(comm.InStringArray(r.Parity, []string{"O", "E"}), r.Parity, "O").(string)
	r.Handler.StopBits = comm.If(r.StopBits == 0, 1, r.StopBits).(int)

	//set timeout
	if r.TimeOut.Seconds() > 0 {
		r.Handler.Timeout = r.TimeOut
	}

	//connect
	if err = r.Handler.Connect(); err != nil {
		return
	}

	//set client
	r.Client = modbus.NewClient(r.Handler)

	//return
	return
}

func (r *RtuClient) Close() (err error) {
	return r.Handler.Close()
}

func (r *RtuClient) ReadHoldingRegisters(address uint16, quantity uint16) (values []int, err error) {
	//读取寄存器
	results, err := r.Client.ReadHoldingRegisters(address, quantity)
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
	return nil, nil
}

func (r *RtuClient) ReadCoils(uint16, uint16) ([]int, error) {
	return nil, nil
}

func (r *RtuClient) ReadInputStatus(uint16, uint16) ([]int, error) {
	return nil, nil
}

func (r *RtuClient) ReadInputRegisters(uint16, uint16) ([]int, error) {
	return nil, nil
}

func (r *RtuClient) WriteSingleRegister(uint16, uint16) error {
	return nil
}
func (r *RtuClient) WriteMultipleRegisters(address, quantity uint16, values []int) error {
	return nil
}
func (r *RtuClient) WriteSingleCoil(uint16, uint16) error {
	return nil
}
func (r *RtuClient) WriteMultipleCoils(uint16, uint16, []int) error {
	return nil
}
