package bacnet

import (
	"fmt"
	"github.com/spf13/cast"
	gobacnet "iotClient/protocol/bacnet/sdk"
	"iotClient/protocol/comm"
	"runtime"
)

func (u *UdpClient) InitBacNet() (err error) {
	//get interfaceByName and port
	interfaceByName := OperateFunc[DefaultInterfaceCode](runtime.GOOS).(string)
	port := OperateFunc[DefaultPortCode](u.Port).(int)

	//client
	client, err := gobacnet.NewClient(interfaceByName, port)
	if err != nil {
		return
	}

	//set client
	u.Client = client

	//return
	return
}

func (u *UdpClient) Close() {
	u.Client.Close()
}

// WhoIs WHOIS
func (u *UdpClient) WhoIs(lowDeviceId, highDeviceId int) (devs []map[string]interface{}, err error) {
	//DevList
	devList, err := u.Client.WhoIs(lowDeviceId, highDeviceId)
	if err != nil {

	}

	//set device list
	var deviceList []map[string]interface{}
	//range device list
	if len(devList) > 0 {
		for _, dev := range devList {
			deviceList = append(deviceList, map[string]interface{}{
				"device_id":        dev.ID.Instance,
				"device_type":      dev.ID.Type,
				"device_type_desc": ObjectTypeDesc[ObjectType(dev.ID.Type)],
				"maxApdu":          dev.MaxApdu,
				"segmentation":     dev.Segmentation,
				"vendor":           dev.Vendor,
				"addr_net":         dev.Addr.Net,
				"addr_len":         dev.Addr.Len,
				"device_address": map[string]string{
					"ip":   fmt.Sprintf("%d.%d.%d.%d", dev.Addr.Mac[0], dev.Addr.Mac[1], dev.Addr.Mac[2], dev.Addr.Mac[3]),
					"port": fmt.Sprintf("%d%d", comm.HexToInt(cast.ToString(dev.Addr.Mac[4])), comm.HexToInt(cast.ToString(dev.Addr.Mac[5]))),
				},
			})
		}
	}

	//set value
	devs = deviceList

	return
}
