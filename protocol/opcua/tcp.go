package opcua

import (
	"context"
	"fmt"
	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"iotClient/protocol/comm"
	"log"
)

/**
https://blog.csdn.net/u013810234/article/details/130189271
https://downloads.prosysopc.com/opc-ua-simulation-server-downloads.php
https://godoc.org/github.com/gopcua/opcua#Option
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
https://github.com/gopcua/opcua/blob/main/examples
*/

func (t *TcpClient) InitOpcUa() (err error) {
	//connect OPC
	opcClient, err := opcua.NewClient(t.EndPoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//set Client
	t.Client = opcClient

	//connect
	if err = t.Client.Connect(context.TODO()); err != nil {
		log.Fatal(err.Error())
		return
	}

	//return
	return
}

func (t *TcpClient) Close() (err error) {
	err = t.Client.Close(context.TODO())
	return
}

func (t *TcpClient) GetPoints() (err error) {
	var resp *ua.GetEndpointsResponse
	resp, err = t.Client.GetEndpoints(context.TODO())
	fmt.Println(resp)
	return
}

func (t *TcpClient) ReadValue(nodeId string) (data map[string]interface{}, err error) {
	//parse node id
	pid, err := ua.ParseNodeID(nodeId)
	if err != nil {
		return
	}

	//read ua
	req := &ua.ReadRequest{
		MaxAge: 2000,
		NodesToRead: []*ua.ReadValueID{
			{NodeID: pid},
		},
		TimestampsToReturn: ua.TimestampsToReturnBoth,
	}

	//read Response
	var resp *ua.ReadResponse
	resp, err = t.Client.Read(context.TODO(), req)
	if err != nil {
		return
	}

	//judge resp is null
	if resp == nil || resp.Results == nil || len(resp.Results) == 0 {
		err = errors.New("Read value is null")
		return
	}

	RetValue := map[string]interface{}{}
	RetValue["EncodingMask"] = resp.Results[0].EncodingMask
	RetValue["Status"] = resp.Results[0].Status
	RetValue["Value"] = map[string]interface{}{
		"mask":                  resp.Results[0].Value.EncodingMask(),
		"arrayLength":           resp.Results[0].Value.ArrayLength(),
		"arrayDimensionsLength": resp.Results[0].Value.ArrayLength(),
		"arrayDimensions":       resp.Results[0].Value.ArrayDimensions(),
		"value":                 resp.Results[0].Value.Value(),
	}
	RetValue["SourceTimestamp"] = resp.Results[0].SourceTimestamp
	RetValue["SourcePicoseconds"] = resp.Results[0].SourcePicoseconds
	RetValue["ServerTimestamp"] = resp.Results[0].ServerTimestamp
	RetValue["ServerPicoseconds"] = resp.Results[0].ServerPicoseconds

	//set value
	data = RetValue
	return
}

// ReadBatchValues 读取批量点位数据
func (t *TcpClient) ReadBatchValues(nodeIds []string) (data []map[string]interface{}, err error) {
	if len(nodeIds) == 0 {
		err = errors.New("nodeId list is null")
		return
	}

	//要返回的数据
	retDatas := []map[string]interface{}{}
	for _, id := range nodeIds {
		var retData map[string]interface{}
		retData, err = t.ReadValue(id)
		if err != nil {
			return
		}
		retDatas = append(retDatas, retData)
	}

	//set value
	data = retDatas

	//return
	return
}

// BrowseNode Node列表
func (t *TcpClient) BrowseNode(nodeId string) (nodeDatas []map[string]interface{}, err error) {
	//parse node id
	pid, err := ua.ParseNodeID(nodeId)
	if err != nil {
		log.Fatal(err)
		return
	}

	//get nodeList
	nodeList, err := browse(context.TODO(), t.Client.Node(pid), "", 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	//set node
	nodeValues := []map[string]interface{}{}
	//range
	for _, def := range nodeList {
		data := comm.StructToMap(def)
		nodeValues = append(nodeValues, data)
	}

	//set values
	nodeDatas = nodeValues

	//return
	return
}

// Write 写入数据
func (t *TcpClient) Write(nodeId, value string) (err error) {
	//parse node id
	pid, err := ua.ParseNodeID(nodeId)
	if err != nil {
		log.Fatalf("parse node id error:%v", err.Error())
		return
	}

	//value
	v, err := ua.NewVariant(value)
	if err != nil {
		log.Fatalf("invalid value: %v", err)
		return
	}

	//request
	req := &ua.WriteRequest{
		NodesToWrite: []*ua.WriteValue{
			{
				NodeID:      pid,
				AttributeID: ua.AttributeIDValue,
				Value: &ua.DataValue{
					EncodingMask: ua.DataValueValue,
					Value:        v,
				},
			},
		},
	}

	//write value
	resp, err := t.Client.Write(context.TODO(), req)
	if err != nil {
		log.Fatalf("Write failed: %s", err)
		return
	}

	//判断CODE
	if len(resp.Results) > 0 {
		errCode := cast.ToInt(resp.Results[0])
		if errCode == 0 {
			return
		}
	}

	//write failed set value
	err = errors.New(fmt.Sprintf("write failed NodeId %s", nodeId))

	//return
	return
}
