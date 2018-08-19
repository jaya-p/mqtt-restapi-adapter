package main

import (
	"fmt"
	//import the Paho Go MQTT library
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Gateway struct {
	GatewayStatus string `json:"gateway_status,omitempty"`
}

func main() {
	// assining variables
	var mqttBrokerUrl string = "tcp://iot.eclipse.org:1883"
	var mqttClientId string = "jinawiLog"
	var mqttTopic string = "jinawi-dev-1"
	var mqttQos byte = 0
	var restApiUrl string = "http://35.198.211.124:1801/gateway/1/toggle"
	var restApiMethod string = "POST"
	var restApiBody string = "{\"gateway_status\": \"off\"}"
	var gateway Gateway

	// set MQTT client variable
	opts := MQTT.NewClientOptions().AddBroker(mqttBrokerUrl)
	opts.SetClientID(mqttClientId)

	// subscribe to a topic
	choke := make(chan [2]string)
	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	mqttClient := MQTT.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("MQTT client is connected to '%s' using '%s' clientId\n", mqttBrokerUrl, mqttClientId)

	if token := mqttClient.Subscribe(mqttTopic, mqttQos, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	if token := mqttClient.Subscribe(mqttTopic, mqttQos, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Printf("MQTT client is subcribed to '%s' topic with '%d' qos\n", mqttTopic, mqttQos)

	// prepare HTTP client
	httpClient := http.Client{}

	// receive MQTT message and call REST API for each changes
	var state = "undefined"
	for true {
		//receive REST API
		req, err := http.NewRequest("GET", "http://35.198.211.124:1801/gateway/1", nil)
		if err != nil {
			panic(err)
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		if 200 != resp.StatusCode {
			panic(body)
		}
		json.Unmarshal(body, &gateway)
		resp.Body.Close()
		message := gateway.GatewayStatus
		// fmt.Println(gateway.GatewayStatus)
		// publish
		if message != state {
			fmt.Printf("RECEIVED API Message: %s \n", message)
			if message == "off" {
				// fmt.Println("Mati")
				token := mqttClient.Publish(mqttTopic, mqttQos, false, "0")
				token.Wait()
				state = message
			} else if message == "on" {
				// fmt.Println("Nyala")
				token := mqttClient.Publish(mqttTopic, mqttQos, false, "1")
				token.Wait()
				state = message
			}
		}

		// fmt.Printf("R\n")

		// receive MQTT message
		select {
		case incoming := <-choke:
			fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
			if incoming[1] == "0" {
				restApiBody = "{\"gateway_status\": \"off\"}"
				state = "off"
			} else if incoming[1] == "1" {
				restApiBody = "{\"gateway_status\": \"on\"}"
				state = "on"
			} else if incoming[1] == "q" {
				fmt.Println(gateway.GatewayStatus)
			} else {
				continue
			}

			// call REST API
			req, err = http.NewRequest(restApiMethod, restApiUrl, strings.NewReader(restApiBody))
			if err != nil {
				panic(err)
			}
			resp, err = httpClient.Do(req)
			if err != nil {
				panic(err)
			}
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			if 200 != resp.StatusCode {
				panic(body)
			}
			fmt.Println(string(body))
			resp.Body.Close()

		default:
			continue
		}

	}

}
