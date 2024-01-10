package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

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
	opts.SetDefaultPublishHandler(m.MessagePubHandler)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(3)
	opts.OnConnect = m.ConnectHandler
	opts.OnConnectionLost = m.ConnectLostHandler

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

// Subscribe 订阅消息
func (m *TcpClient) Subscribe(topic string, subFunc SubscribeHandler) {
	if token := m.Client.Subscribe(topic, 0, mqtt.MessageHandler(subFunc)); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}

// DisConnect 断开链接
func (m *TcpClient) DisConnect() {
	if m.Client.IsConnected() {
		m.Client.Disconnect(30)
	}
}
