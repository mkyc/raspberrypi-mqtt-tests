package main

import (
	"github.com/mkyc/raspberrypi-mqtt-tests/pkg/state"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/mqtt"
	"log"
)

const (
	address  = "tcp://0.0.0.0:1883"
	clientID = "c2"
	topic    = "pir"
	name     = "sender"
)

func main() {
	mqttAdaptor := mqtt.NewAdaptor(address, clientID)

	work := func() {
		mqttAdaptor.On(topic, func(msg mqtt.Message) {
			log.Println(msg.Payload())
			s := state.State{}
			err := s.Deserialize(msg.Payload())
			if err != nil {
				log.Fatal(err)
			}
			log.Println(s)
		})
	}

	robot := gobot.NewRobot(name,
		[]gobot.Connection{mqttAdaptor},
		work,
	)

	robot.Start()
}
