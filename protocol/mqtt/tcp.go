package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// var MessagePubHandler pubHandler
// var ConnectHandler connectHandler
//var ConnectLostHandler lostHandler

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("hello1")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Println("hello2")
}

// InitMqtt init mqtt
func (m *TcpClient) InitMqtt() (err error) {
	//New Client Options
	opts := mqtt.NewClientOptions()
	//add
	opts.AddBroker(m.Broker)
	opts.SetClientID(m.ClientId)

	//set username password
	if m.UserName != "" && m.Password != "" {
		opts.SetUsername(m.UserName)
		opts.SetPassword(m.Password)
	}
	opts.SetDefaultPublishHandler(MessagePubHandler)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(3)
	opts.OnConnect = ConnectHandler
	opts.OnConnectionLost = ConnectLostHandler

	//set client
	m.Client = mqtt.NewClient(opts)
	if token := m.Client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	//return
	return
}

// Publish 发布消息
func (m *TcpClient) Publish(topic, payLoad string) (err error) {
	if token := m.Client.Publish(topic, 0, true, payLoad); token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	return
}

// DisConnect 断开链接
func (m *TcpClient) DisConnect() {
	if m.Client.IsConnected() {
		m.Client.Disconnect(30)
	}
}
