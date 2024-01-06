package opcua

import (
	"github.com/gopcua/opcua"
)

type OpcUaClient interface {
	InitOpcUa() error
	Close() error
	ReadValue(nodeId string) (map[string]interface{}, error)
}

type TcpClient struct {
	EndPoint string `json:"end_point"` //TCP 地址  opc.tcp://1:53530/OPCUA/SimulationServer
	Client   *opcua.Client
}
