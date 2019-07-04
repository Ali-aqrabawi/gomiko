package examples

import (
	"fmt"
	"github.com/Ali-aqrabawi/gomiko/pkg"
	"log"
)

func mainASA() {
	// example of device that does not have Secret
	exampleNoSecretASA()
	// example of device that has Secret
	exampleWithSecrectASA()
}

// create Cisco IOS Device without Secret and connect to it.
// then execute some commands
func exampleNoSecretASA() {
	//create Cisco IOS Device without Secret and connect to it.
	device, err := gomiko.NewDevice("192.168.1.102", "admin", "mySecret", "cisco_asa", 22)
	if err != nil {
		log.Fatal(err)
	}

	//Open Session with device
	if err := device.Connect(); err != nil {
		log.Fatal(err)
	}

	// send command
	output1, _ := device.SendCommand("show nat")

	// send a set of config commands
	commands := []string{"object network thisObject", "host 44.6.3.1"}
	output2, _ := device.SendConfigSet(commands)

	device.Disconnect()

	fmt.Println(output1)
	fmt.Println(output2)
}

// create Cisco IOS Device without Secret and connect to it.
// then execute some commands
func exampleWithSecrectASA() {
	// create Cisco IOS Device with SecretOption and connect to it.
	secOption := gomiko.SecretOption("mySecret")
	device, err := gomiko.NewDevice(
		"192.168.1.102", "admin",
		"mySecret", "cisco_asa",
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
	output1, _ := device.SendCommand("show nat")

	// send a set of config commands
	commands := []string{"object network thisObject", "host 44.6.3.1"}
	output2, _ := device.SendConfigSet(commands)

	device.Disconnect()

	fmt.Println(output1)
	fmt.Println(output2)

}
