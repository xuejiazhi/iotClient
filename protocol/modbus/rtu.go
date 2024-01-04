package modbus

import (
	"errors"
	"github.com/goburrow/modbus"
	"iotClient/protocol/comm"
	"log"
	"runtime"
	"strconv"
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

// ReadHoldingRegisters 读取寄存器数据
// address 地址
// quantity 数量
func (r *RtuClient) ReadHoldingRegisters(address uint16, quantity uint16) (values []int, err error) {
	//读取寄存器
	return ModbusOperate["readHoldingRegisters"](RegClient{Address: address, Quantity: quantity, Client: r.Client})
}

// ReadCoils 读取线圈
// address 地址
// quantity 数量
func (r *RtuClient) ReadCoils(address, quantity uint16) (values []int, err error) {
	//读取线圈
	return ModbusOperate["readCoils"](RegClient{Address: address, Quantity: quantity, Client: r.Client})
}

// ReadInputStatus 输入状态
// address 地址
// quantity 数量
func (r *RtuClient) ReadInputStatus(address, quantity uint16) ([]int, error) {
	//读取输出
	return ModbusOperate["readInputStatus"](RegClient{Address: address, Quantity: quantity, Client: r.Client})
}

// ReadInputRegisters 输入寄存器
// address 地址
// quantity 数量
func (r *RtuClient) ReadInputRegisters(address, quantity uint16) ([]int, error) {
	return ModbusOperate["readInputRegisters"](RegClient{Address: address, Quantity: quantity, Client: r.Client})
}

// WriteSingleRegister 写入单个寄存器
// address 地址
// value 值
func (r *RtuClient) WriteSingleRegister(address, value uint16) (err error) {
	//写入单个寄存器
	result, err := r.Client.WriteSingleRegister(address, value)
	if err != nil {
		return
	}

	//check less len
	if len(result) < 2 {
		err = errors.New("less than 2 Byte")
		return
	}

	//return
	return
}

// WriteMultipleRegisters 批量写入寄存器
// address 地址
// quantity 数量
// values 更新的值列表
func (r *RtuClient) WriteMultipleRegisters(address, quantity uint16, values []int) (err error) {
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
	_, err = r.Client.WriteMultipleRegisters(address, quantity, dataBytes)

	//return
	return
}

// WriteSingleCoil 写入单个线圈
func (r *RtuClient) WriteSingleCoil(address, value uint16) (err error) {
	//是否在数组里面
	if !comm.InIntArray(int(value), []int{0, 1}) {
		err = errors.New("modbus: state '1' must be either 1 (ON) or 0 (OFF)")
		return
	}

	//取值
	coilValue := comm.If(value == 1, CoilStateOn, CoilStateOff).(int)
	//写入值
	_, err = r.Client.WriteSingleCoil(address, uint16(coilValue))
	if err != nil {
		return
	}
	return
}

// WriteMultipleCoils 批量写入线圈
func (r *RtuClient) WriteMultipleCoils(address, quantity uint16, values []int) (err error) {
	//check eq len
	if c := GetOperate["checkEqLen"](len(values), int(quantity)); c != nil {
		err = c.(error)
		return
	}

	//转化值
	decimal, err := func(x []int) (decimal int64, err error) {
		binaryStr := ""
		for _, v := range x {
			binaryStr = strconv.Itoa(v) + binaryStr
		}
		return strconv.ParseInt(binaryStr, 2, 64)
	}(values)

	//err return
	if err != nil {
		log.Fatal("write multiple coils error,convert decimal error,", err.Error())
		return
	}

	//write value
	_, err = r.Client.WriteMultipleCoils(address, quantity, []byte{uint8(decimal)})
	if err != nil {
		log.Fatal("write multiple coils error,", err.Error())
		return
	}

	return
}
