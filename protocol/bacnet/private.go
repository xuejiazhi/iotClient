package bacnet

import (
	"iotClient/protocol/comm"
)

type ObjectType uint16

const (
	AnalogInput       ObjectType = 0
	AnalogOutput      ObjectType = 1
	AnalogValue       ObjectType = 2
	BinaryInput       ObjectType = 3
	BinaryOutput      ObjectType = 4
	BinaryValue       ObjectType = 5
	DeviceType        ObjectType = 8
	File              ObjectType = 10
	MultiStateInput   ObjectType = 13
	NotificationClass ObjectType = 15
	MultiStateValue   ObjectType = 19
	TrendLog          ObjectType = 20
	CharacterString   ObjectType = 40
)

var ObjectTypeDesc = map[ObjectType]string{
	AnalogInput:       "AnalogInput",
	AnalogOutput:      "AnalogOutput",
	AnalogValue:       "AnalogValue",
	BinaryInput:       "BinaryInput",
	BinaryOutput:      "BinaryOutput",
	BinaryValue:       "BinaryValue",
	DeviceType:        "DeviceType",
	File:              "File",
	MultiStateInput:   "MultiStateInput",
	NotificationClass: "NotificationClass",
	MultiStateValue:   "MultiStateValue",
	TrendLog:          "TrendLog",
	CharacterString:   "CharacterString",
}

type Operate func(x any) any

const (
	DefaultPortCode      = 2000001
	DefaultInterfaceCode = 2000002

	LinuxInterfaceName   = "eth0"
	WindowsInterfaceName = "以太网"
)

func GetDefaultPort(x any) any { return comm.If(x.(int) == 0, DefaultBacNetPort, x) }
func GetDefaultInterfaceName(x any) any {
	//默认的interfaceName
	interfaceName := LinuxInterfaceName
	switch x.(string) {
	case "windows":
		interfaceName = WindowsInterfaceName
	case "linux":
		interfaceName = LinuxInterfaceName
	default:
		interfaceName = LinuxInterfaceName
	}
	//返回
	return interfaceName
}

var OperateFunc = map[int]Operate{
	DefaultPortCode:      GetDefaultPort,
	DefaultInterfaceCode: GetDefaultInterfaceName,
}
