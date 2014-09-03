package main

import (
	"../."
	"encoding/json"
	"fmt"
	"os"
)

const (
	ClientID          = "<client-id>"
	State             = "STATE"
	ClientSecret      = "<client-secret>"
	AuthorizationCode = "<authorization-code> - https://developer.nest.com/documentation/how-to-auth"
	Token             = "<token>"
)

func main() {
	client := nest.New(ClientID, State, ClientSecret, AuthorizationCode)
	client.Token = Token
	devicesChan := make(chan *nest.Devices)
	go func() {
		client.DevicesStream(func(devices *nest.Devices, err error) {
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			devicesChan <- devices
		})
	}()

	for i := 0; i < 7; i++ {
		devices := <-devicesChan
		thermostat := devices.Thermostats["1cf6CGENM20W3UsKiJTT4Cqpa4SSjzbd"]
		switch i {
		case 0:
			logEvent(devices, i)
			fmt.Println("Setting target temp")
			err := thermostat.SetTargetTempF(thermostat.TargetTemperatureF + 1)
			if err != nil {
				fmt.Printf("Error: %s - %d\n", err.Description, i)
				os.Exit(2)
			}
		case 1:
			logEvent(devices, i)
			fmt.Println("Setting HvacMode to HeatCool")
			err := thermostat.SetHvacMode(nest.HeatCool)
			if err != nil {
				fmt.Printf("Error: %s - %d\n", err.Description, i)
				os.Exit(2)
			}
		case 2:
			logEvent(devices, i)
			fmt.Println("Setting TargetTempHighLow")
			err := thermostat.SetTargetTempHighLowF(thermostat.TargetTemperatureHighF+1, thermostat.TargetTemperatureLowF+1)
			if err != nil {
				fmt.Printf("Error: %s - %d\n", err.Description, i)
				os.Exit(2)
			}
		case 3:
			logEvent(devices, i)
			fmt.Println("Setting HvacMode to Heat")
			err := thermostat.SetHvacMode(nest.Heat)
			if err != nil {
				fmt.Printf("Error: %s - %d\n", err.Description, i)
				os.Exit(2)
			}
		case 4:
			logEvent(devices, i)
			fmt.Println("Setting FanTimeActive to true")
			err := thermostat.SetFanTimerActive(true)
			if err != nil {
				fmt.Printf("Error: %s - %d\n", err.Description, i)
				os.Exit(2)
			}
		case 5:
			logEvent(devices, i)
			fmt.Println("Setting FanTimeActive to false")
			err := thermostat.SetFanTimerActive(false)
			if err != nil {
				fmt.Printf("Error: %s - %d\n", err.Description, i)
				os.Exit(2)
			}
		case 6:
			logEvent(devices, i)
			break
		}
	}
}

func logEvent(devices *nest.Devices, cnt int) {
	fmt.Printf(">>>>>%d<<<<<\n", cnt)
	data, _ := json.MarshalIndent(devices, "", "  ")
	fmt.Println(string(data))
	fmt.Printf(">>>>>%d<<<<<\n", cnt)
}
