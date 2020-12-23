package main

import (
	"github.com/mkyc/raspberrypi-mqtt-tests/pkg/state"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/keyboard"
	"gobot.io/x/gobot/platforms/mqtt"
	"log"
	"time"
)

const (
	address  = "tcp://0.0.0.0:1883"
	clientID = "c1"
	topic    = "pir"
	name     = "sender"
)

func main() {
	s := state.State{}

	mqttAdaptor := mqtt.NewAdaptor(address, clientID)
	keys := keyboard.NewDriver()

	work := func() {
		keys.On(keyboard.Key, func(data interface{}) {
			event := data.(keyboard.KeyEvent)
			switch event.Key {
			case keyboard.One:
				s.PIR1 = !s.PIR1
			case keyboard.Two:
				s.PIR2 = !s.PIR2
			case keyboard.Three:
				s.PIR3 = !s.PIR3
			case keyboard.Four:
				s.PIR4 = !s.PIR4
			}
		})

		gobot.Every(1*time.Second, func() {
			log.Println(s)
			b, err := s.Serialize()
			if err != nil {
				log.Fatal(err)
			}
			mqttAdaptor.Publish(topic, b)
			log.Println("sent")
		})
	}

	robot := gobot.NewRobot(name,
		[]gobot.Connection{mqttAdaptor},
		[]gobot.Device{keys},
		work,
	)

	robot.Start()
}
