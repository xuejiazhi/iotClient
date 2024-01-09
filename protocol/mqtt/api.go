package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type MqttClient interface {
	InitMqtt() error
	DisConnect()
	Publish(string, string) error
}

type TcpClient struct {
	Broker             string                     `json:"broker"`
	ClientId           string                     `json:"client_id"`
	UserName           string                     `json:"user_name"`
	Password           string                     `json:"password"`
	Client             mqtt.Client                `json:"client"`
	MessagePubHandler  mqtt.MessageHandler        `json:"messagePubHandler"`
	ConnectHandler     mqtt.OnConnectHandler      `json:"connectHandler"`
	ConnectLostHandler mqtt.ConnectionLostHandler `json:"connectLostHandler"`
}
