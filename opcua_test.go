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

// 读取单个点位数据
func Test_ReadOpcUa(t *testing.T) {
	opc := initOpcUa()
	value, err := opc.ReadValue("ns=3;i=1008")
	if err != nil {
		log.Fatal(err.Error())
	}
	for key, value := range value {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
	opc.Close()
}

func Test_GetEndPoints(t *testing.T) {
	opc := initOpcUa()
	err := opc.GetPoints()
	if err != nil {
		log.Fatal(err.Error())
	}
	opc.Close()
}

// 批量读取点位数据
func Test_ReadOpeUaValues(t *testing.T) {
	opc := initOpcUa()
	value, err := opc.ReadBatchValues([]string{"ns=3;i=1003", "ns=3;i=1004", "ns=3;i=1008"})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(value)
	opc.Close()
}

func Test_BrowseNode(t *testing.T) {
	opc := initOpcUa()
	nodeList, err := opc.BrowseNode("ns=3;s=85/0:Simulation")
	log.Print(nodeList)
	log.Print(err)
}
