package examples

import (
	"fmt"
	"github.com/Ali-aqrabawi/gomiko/pkg"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

func mainASA() {
	// example of device that does not have Secret
	exampleNoSecretASA()
	// example of device that has Secret
	exampleWithSecrectASA()
	// example of devoce that created by passing *ssh.Client.
	exampleSSHClientASA()
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
	if err := device.OpenSession(); err != nil {
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
	if err := device.OpenSession(); err != nil {
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

// create Cisco IOS Device using *ssh.Client connection
// then execute some commands,
// NOTE: you should call .Dial() before passing the client to .NewDeviceFromClient()
func exampleSSHClientASA() {

	// create secret option
	opt := gomiko.SecretOption("mySecret")

	// create ssh client config
	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("mySecret")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout: 6 * time.Second,
	}
	var ciphers = []string{"3des-cbc", "aes128-cbc", "aes192-cbc", "aes256-cbc", "aes128-ctr"}
	sshConfig.Ciphers = append(sshConfig.Ciphers, ciphers...)

	// connect to device
	conn, _ := ssh.Dial("tcp", "192.168.1.102:22", sshConfig)

	// create device using that connection which will be used by gomiko to start session
	device, err := gomiko.NewDeviceFromClient(conn, "cisco_asa", opt)

	if err != nil {
		log.Fatal(err)
	}
	// Open Session to device.
	if err := device.OpenSession(); err != nil {
		log.Fatal(err)
	}

	// execut commands
	output1, _ := device.SendCommand("show nat")

	// send a set of config commands
	commands := []string{"object network thisObject", "host 44.6.3.1"}
	result2, _ := device.SendConfigSet(commands)

	fmt.Println(output1)
	fmt.Println(result2)
}
