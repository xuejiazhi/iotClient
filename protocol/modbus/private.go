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
	return comm.If(len(x.([]byte)) > y.(int), nil, errors.New(fmt.Sprintf("less than %v", y)))
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
