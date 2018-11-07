package driver

import (
	"errors"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"net/url"
	"time"
)

func ConnectMQTT(clientId string, uri *url.URL) (mqtt.Client, error) {
	if clientId == "" {
		return nil, errors.New("couldn't connect to the MQTT Broker, invalid client id")
	}

	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		return nil, errors.New("couldn't connect to the MQTT Broker, no connection")
	}
	return client, nil
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}