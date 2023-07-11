package examples

import (
	"fmt"
	"github.com/Ali-aqrabawi/gomiko/pkg"
	"log"
)

func mainIOS() {
	// example of device that does not have Secret
	exampleNoSecret()
	// example of device that has Secret
	exampleWithSecrect()
	// example of device with custom timeout that has Secret
	exampleWithTimeoutWithSecrect()
}

// create Cisco IOS Device without Secret and connect to it.
// then execute some commands
func exampleNoSecret() {
	//create Cisco IOS Device without Secret and connect to it.
	device, err := gomiko.NewDevice("192.168.1.1", "admin", "mySecret", "cisco_ios", 22)
	if err != nil {
		log.Fatal(err)
	}

	//Open Session with device
	if err := device.Connect(); err != nil {
		log.Fatal(err)
	}

	// send command
	output1, _ := device.SendCommand("show vlan")

	// send a set of config commands
	commands := []string{"vlan 120", "name v120"}
	output2, _ := device.SendConfigSet(commands)

	device.Disconnect()

	fmt.Println(output1)
	fmt.Println(output2)
}

// create Cisco IOS Device with Secret and connect to it.
// then execute some commands
func exampleWithSecrect() {
	// create Cisco IOS Device with SecretOption and connect to it.
	secOption := gomiko.SecretOption("mySecret")
	device, err := gomiko.NewDevice(
		"192.168.1.1", "admin",
		"mySecret", "cisco_ios",
		22, secOption,
	)
	if err != nil {
		log.Fatal(err)
	}

	//Open Session with device
	if err := device.Connect(); err != nil {
		log.Fatal(err)
	}

	// send command
	output1, _ := device.SendCommand("show vlan")

	// send a set of config commands
	commands := []string{"vlan 120", "name v120"}
	output2, _ := device.SendConfigSet(commands)

	device.Disconnect()

	fmt.Println(output1)
	fmt.Println(output2)

}

// create Cisco IOS Device with custom Timeout and with Secret and connect to it.
// then execute some commands
func exampleWithTimeoutWithSecrect() {
	// create Cisco IOS Device with SecretOption and connect to it.
	secOption := gomiko.SecretOption("mySecret")
	device, err := gomiko.NewDeviceWithTimeout(
		"192.168.1.1", "admin",
		"mySecret", "cisco_ios",
		22, 10, secOption,
	)
	if err != nil {
		log.Fatal(err)
	}

	//Open Session with device
	if err := device.Connect(); err != nil {
		log.Fatal(err)
	}

	// send command
	output1, _ := device.SendCommand("show vlan")

	// send a set of config commands
	commands := []string{"vlan 120", "name v120"}
	output2, _ := device.SendConfigSet(commands)

	device.Disconnect()

	fmt.Println(output1)
	fmt.Println(output2)

}
