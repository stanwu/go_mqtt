package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/shirou/gopsutil/v3/cpu"
)

const (
	broker   = "tcp://broker.emqx.io:1883"
	topic    = "test/topic123"
	clientID = "golang-123"
)

func getCPULoad() {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Printf("Error retrieving CPU load: %v\n", err)
		return
	}
	fmt.Printf("CPU Load: %.2f%%\n", percentages[0])
}

func main() {
	// Create MQTT client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)

	// Create MQTT client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Error connecting to broker: %v\n", token.Error())
		return
	}
	defer client.Disconnect(250)

	// Subscribe to the topic
	if token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}); token.Wait() && token.Error() != nil {
		fmt.Printf("Error subscribing to topic: %v\n", token.Error())
		return
	}

	// Publish messages every 5 seconds
	go func() {
		for {
			message := map[string]interface{}{}
			percentages, err := cpu.Percent(0, false)
			if err != nil {
				fmt.Printf("Error retrieving CPU load: %v\n", err)
			} else {
				message["CPU"] = fmt.Sprintf("%.2f", percentages[0])
			}
			payload, err := json.Marshal(message)
			if err != nil {
				fmt.Printf("Error marshalling JSON: %v\n", err)
				continue
			}
			if token := client.Publish(topic, 1, false, payload); token.Wait() && token.Error() != nil {
				fmt.Printf("Error publishing message: %v\n", token.Error())
			}
			time.Sleep(5 * time.Second)
		}
	}()

	// Wait for termination signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Exiting...")
}
