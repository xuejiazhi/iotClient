package modbus

import (
	"encoding/binary"
	"errors"
	"fmt"
	"iotClient/protocol/comm"
	"strconv"
)

func getRegisterValue(dataBytes []byte) int {
	var retData int
	if len(dataBytes) == 2 {
		retData = int(dataBytes[0])*256 + int(dataBytes[1])
	}
	return retData
}

// intToBytes 整形转换成字节
func intToBytes(ns []int) []byte {
	// 创建一个字节数组
	b := make([]byte, 0, 2*binary.MaxVarintLen64)
	for _, n := range ns {
		n1 := n / 256
		n2 := n % 256
		b = append(b, byte(n1), byte(n2))
	}
	return b
}

// 将二进制转成十进制
func intoCoilsByte(ns []int) (int64, error) {
	binaryStr := ""
	for _, n := range ns {
		binaryStr += strconv.Itoa(n)
	}
	//parse
	return strconv.ParseInt(binaryStr, 2, 64)
}

// Operate function
type Operate func(x, y any) any

func GetModBusSlaveID(x, y any) any { return comm.If(x.(uint8) > 0, x, y).(uint8) }
func GetModBusTcpAddr(x, y any) any { return comm.If(x != "", x, y).(string) }
func GetModBusTcpPort(x, y any) any { return comm.If(x.(int) > 0, x, y).(int) }
func CheckLessLen(x, y any) any {
	return comm.If(len(x.([]byte)) == y.(int), nil, errors.New(fmt.Sprintf("less than %v", y)).(error))
}
func CheckEqualLen(x, y any) any {
	eq := comm.If(x.(int) == y.(int), nil, errors.New(fmt.Sprintf("length is not equal, %v", y)))
	if eq != nil {
		return eq.(error)
	} else {
		return eq
	}
}

var GetOperate = map[string]Operate{
	"slaveId":      GetModBusSlaveID,
	"tcpAddr":      GetModBusTcpAddr,
	"tcpPort":      GetModBusTcpPort,
	"checkLessLen": CheckLessLen,
	"checkEqLen":   CheckEqualLen,
}

type OperateModbus func(rc RegClient) ([]int, error)

var ModbusOperate = map[string]OperateModbus{
	"readRegister": func(rc RegClient) (values []int, err error) {
		//读取寄存器
		results, err := GetResult(rc, "readRegister", 2)

		//设置数据
		for i := 0; i < len(results); i = i + 2 {
			//一个数据为两个byte
			dataBytes := results[i : i+2]
			if len(dataBytes) == 2 {
				values = append(values, getRegisterValue(dataBytes))
			}
		}
		return
	},
	"readCoils": func(rc RegClient) (values []int, err error) {
		// 读点位线圈数据
		results, err := GetResult(rc, "readCoils", 1)
		if err != nil {
			return
		}
		//取data
		values = comm.DecimalToBinary(int(results[0]))
		return
	},
	"readInputs": func(rc RegClient) (values []int, err error) {
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
	},
}

func GetResult(rc RegClient, opt string, lessThan int) (results []byte, err error) {
	switch opt {
	case "readCoils":
		results, err = rc.Client.ReadCoils(rc.Address, rc.Quantity)
	case "readRegister":
		results, err = rc.Client.ReadHoldingRegisters(rc.Address, rc.Quantity)
	case "readInputs":
		results, err = rc.Client.ReadDiscreteInputs(rc.Address, rc.Quantity)
	default:
		err = errors.New("Opt is error")
		return
	}

	//check less len
	if len(results) < lessThan {
		err = errors.New(fmt.Sprintf("length is not than %d", lessThan))
		return
	}
	//return
	return
}
