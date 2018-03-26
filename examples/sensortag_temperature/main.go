package main

import (
	log "github.com/sirupsen/logrus"

	"git.enexoma.de/r/smartcontrol/libraries/go-bluetooth/api"
	"git.enexoma.de/r/smartcontrol/libraries/go-bluetooth/devices"
)

// example of reading temperature from a TI sensortag
func main() {

	var tagAddress = "B0:B4:48:C9:4B:01"

	dev, err := api.GetDeviceByAddress(tagAddress)
	if err != nil {
		panic(err)
	}

	if dev == nil {
		panic("Device not found")
	}

	err = dev.Connect()
	if err != nil {
		panic(err)
	}

	sensorTag, err := devices.NewSensorTag(dev)
	if err != nil {
		panic(err)
	}

	// var readTemperature = func() {
	// 	temp, err := sensorTag.Temperature.Read()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Printf("Temperature %v°", temp)
	// }

	var notifyTemperature = func(fn func(temperature float64)) {
		sensorTag.Temperature.StartNotify()
		select {}
	}

	// readTemperature()
	notifyTemperature(func(t float64) {
		log.Infof("Temperature update: %f", t)
	})

}
