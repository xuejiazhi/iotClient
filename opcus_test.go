package main

import (
	"fmt"
	op2 "iotClient/protocol/opcua"
	"log"
	"testing"
)

func initOpcUa() op2.OpcUaClient {
	var opcSer op2.OpcUaClient = &op2.TcpClient{
		EndPoint: "opc.tcp://1:53530/OPCUA/SimulationServer",
	}
	opcSer.InitOpcUa()
	return opcSer
}

func Test_connectOpcUa(t *testing.T) {
	var opcSer op2.OpcUaClient = &op2.TcpClient{
		EndPoint: "opc.tcp://1:53530/OPCUA/SimulationServer",
	}
	opcSer.InitOpcUa()
	opcSer.Close()
}

func Test_ReadOpcUa(t *testing.T) {
	opc := initOpcUa()
	value, err := opc.ReadValue("ns=3;i=1008")
	if err != nil {
		log.Fatal(err.Error())
	}
	for key, value := range value {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}
