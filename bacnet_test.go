package main

import (
	"iotClient/protocol/bacnet"
	"log"
	"testing"
)

func initBacnet() bacnet.BacnetClient {
	var bacnet bacnet.BacnetClient = &bacnet.UdpClient{
		InterfaceByName: "以太网",
		Port:            47808,
	}
	_ = bacnet.InitBacNet()
	return bacnet
}

func Test_WhoIs(t *testing.T) {
	b := initBacnet()
	defer b.Close()
	devlist, _ := b.WhoIs(2359386, 2359388)
	log.Fatal(devlist)
}
