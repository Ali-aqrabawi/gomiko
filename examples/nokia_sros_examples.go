package main

import (
	"fmt"
	"github.com/Ali-aqrabawi/gomiko/pkg"
	"log"
)

func main() {
	exampleBasicSROS()
}

func exampleBasicSROS() {
	device, err := gomiko.NewDevice("192.168.1.1", "admin", "admin", "nokia_sros", 22)
	if err != nil {
		log.Fatal(err)
	}

	//Open Session with device
	if err := device.Connect(); err != nil {
		log.Fatal(err)
	}

	// send command
	output1, _ := device.SendCommand("show port")
	output2, _ := device.SendCommand("show uptime")

	// send config command for classic CLI
	_, _ = device.SendCommand("/configure system name TEST")
	_, _ = device.SendCommand("admin save")

	// send a set of config commands
	// NOTE: only works for MD-CLI
	// NOTE: gomiko will automatically commit the changes,so no need to pass "commit"
	commands := []string{"system name TEST"}
	output3, _ := device.SendConfigSet(commands)

	device.Disconnect()

	fmt.Println(output1)
	fmt.Println(output2)
	fmt.Println(output3)
}
