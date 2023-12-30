package modbus

import (
	"errors"
	"fmt"
	"iotClient/protocol/comm"
)

func getRegisterValue(dataBytes []byte) int {
	var retData int
	if len(dataBytes) == 2 {
		retData = int(dataBytes[0])*256 + int(dataBytes[1])
	}
	return retData
}

// Operate function
type Operate func(x, y any) any

func GetModBusSlaveID(x, y any) any { return comm.If(x.(uint8) > 0, x, y).(uint8) }
func GetModBusTcpAddr(x, y any) any { return comm.If(x != "", x, y).(string) }
func GetModBusTcpPort(x, y any) any { return comm.If(x.(int) > 0, x, y).(int) }
func CheckLessLen(x, y any) any {
	return comm.If(len(x.([]byte)) > y.(int), nil, errors.New(fmt.Sprintf("less than %v", y)))
}

var GetOperate = map[string]Operate{
	"slaveId":      GetModBusSlaveID,
	"tcpAddr":      GetModBusTcpAddr,
	"tcpPort":      GetModBusTcpPort,
	"checkLessLen": CheckLessLen,
}
