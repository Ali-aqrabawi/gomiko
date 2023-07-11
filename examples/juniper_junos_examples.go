package examples

import (
	"fmt"
	"github.com/Ali-aqrabawi/gomiko/pkg"
	"log"
)

func mainJunos() {
	// example of device that has Secret
	exampleBasicJunos()
	// example of device with custom timeout
	exampleBasicJunosWithTimeout()
}

// create Junos Device with Secret and connect to it.
// then execute some commands
func exampleBasicJunos() {
	//create juniper device
	device, err := gomiko.NewDevice("192.168.1.102", "admin", "mySecret", "juniper", 22)
	if err != nil {
		log.Fatal(err)
	}

	//Open Session with device
	if err := device.Connect(); err != nil {
		log.Fatal(err)
	}

	// send command
	output1, _ := device.SendCommand("show interfaces brief")

	// send a set of config commands
	// NOTE: gomiko will automatically commit the changes,so no need to pass "commit"
	commands := []string{"set routing-options static route 192.168.47.0/24 next-hop 172.16.1.2"}
	output2, _ := device.SendConfigSet(commands)

	device.Disconnect()

	fmt.Println(output1)
	fmt.Println(output2)
}

// create Junos Device with Secret and connect to it.
// then execute some commands
func exampleBasicJunosWithTimeout() {
	//create juniper device
	device, err := gomiko.NewDeviceWithTimeout("192.168.1.102", "admin", "mySecret", "juniper", 22, 10)
	if err != nil {
		log.Fatal(err)
	}

	//Open Session with device
	if err := device.Connect(); err != nil {
		log.Fatal(err)
	}

	// send command
	output1, _ := device.SendCommand("show interfaces brief")

	// send a set of config commands
	// NOTE: gomiko will automatically commit the changes,so no need to pass "commit"
	commands := []string{"set routing-options static route 192.168.47.0/24 next-hop 172.16.1.2"}
	output2, _ := device.SendConfigSet(commands)

	device.Disconnect()

	fmt.Println(output1)
	fmt.Println(output2)
}
