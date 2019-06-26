package examples

import (
	"fmt"
	"github.com/Ali-aqrabawi/gomiko/pkg"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

func mainIOS() {
	// example of device that does not have Secret
	exampleNoSecret()
	// example of device that has Secret
	exampleWithSecrect()
	// example of devoce that created by passing *ssh.Client.
	exampleSSHClient()
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
	if err := device.OpenSession(); err != nil {
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

// create Cisco IOS Device without Secret and connect to it.
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
	if err := device.OpenSession(); err != nil {
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

// create Cisco IOS Device using *ssh.Client connection
// then execute some commands,
// NOTE: you should call .Dial() before passing the client to .NewDeviceFromClient()
func exampleSSHClient() {

	// create secret option
	opt := gomiko.SecretOption("mySecret")

	// create ssh client config
	sshConfig := &ssh.ClientConfig{User: "admin", Auth: []ssh.AuthMethod{ssh.Password("mySecret")}, HostKeyCallback: ssh.InsecureIgnoreHostKey(), Timeout: 6 * time.Second}
	var ciphers = []string{"3des-cbc", "aes128-cbc", "aes192-cbc", "aes256-cbc", "aes128-ctr"}
	sshConfig.Ciphers = append(sshConfig.Ciphers, ciphers...)

	// connect to device
	conn, _ := ssh.Dial("tcp", "192.168.1.1:22", sshConfig)

	// create device using that connection which will be used by gomiko to start session
	device,err := gomiko.NewDeviceFromClient(conn,"cisco_ios",opt)

	if err != nil{
		log.Fatal(err)
	}
	// Open Session to device.
	if err :=device.OpenSession(); err!=nil{
		log.Fatal(err)
	}

	// execut commands
	result, _ := device.SendCommand("show vlan")
	cmds := []string{"vlan 898","name v898"}
	result2, _ := device.SendConfigSet(cmds)

	fmt.Println(result)
	fmt.Println(result2)
}
