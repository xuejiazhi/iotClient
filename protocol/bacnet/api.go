package bacnet

import gobacnet "iotClient/protocol/bacnet/sdk"

const (
	DefaultBacNetPort = 0xBAC0
)

type BacnetClient interface {
	InitBacNet() error
	WhoIs(lowDeviceId, highDeviceId int) ([]map[string]interface{}, error)
	Close()
}

type UdpClient struct {
	InterfaceByName string           `json:"interface_by_name"` //Bacnet Interface Name
	Port            int              `json:"port"`              //端口
	Client          *gobacnet.Client `json:"client"`            //bacnet client
}
