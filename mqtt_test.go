package main

import (
	"fmt"
	mqtt2 "github.com/eclipse/paho.mqtt.golang"
	mqtt "iotClient/protocol/mqtt"
	"testing"
	"time"
)

func initMqtt() mqtt.MqttClient {
	var mqttSer mqtt.MqttClient = &mqtt.TcpClient{
		Broker:            "tcp://192.168.31.201:1883",
		ClientId:          "mqttx_123456",
		MessagePubHandler: MessageH,
	}
	_ = mqttSer.InitMqtt()
	return mqttSer
}

func Test_initMqtt(t *testing.T) {
	m := initMqtt()
	for i := 0; i < 100; i++ {
		_ = m.Publish("xx-topic", fmt.Sprintf("test123456_%v", i))
		time.Sleep(1 * time.Second)
	}

	select {}
}

func MessageH(client mqtt2.Client, message mqtt2.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}
