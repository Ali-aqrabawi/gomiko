package main

import (
	"fmt"
	"github.com/Ali-aqrabawi/gomiko/pkg"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

func mainJunos() {
	// example of device that has Secret
	exampleBasicJunos()
	// example of devoce that created by passing *ssh.Client.
	exampleSSHClientJunos()
}

// create Cisco IOS Device without Secret and connect to it.
// then execute some commands
func exampleBasicJunos() {
	//create juniper device
	device, err := gomiko.NewDevice("192.168.1.102", "admin", "mySecret", "juniper", 22)
	if err != nil {
		log.Fatal(err)
	}

	//Open Session with device
	if err := device.OpenSession(); err != nil {
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



// create juniper Device using *ssh.Client connection
// then execute some commands,
// NOTE: you should call .Dial() before passing the client to .NewDeviceFromClient()
func exampleSSHClientJunos() {

	// create secret option

	// create ssh client config
	sshConfig := &ssh.ClientConfig{
		User:            "admin",
		Auth:            []ssh.AuthMethod{ssh.Password("password")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         6 * time.Second,
	}
	var ciphers = []string{"3des-cbc", "aes128-cbc", "aes192-cbc", "aes256-cbc", "aes128-ctr"}
	sshConfig.Ciphers = append(sshConfig.Ciphers, ciphers...)

	// connect to device
	conn, _ := ssh.Dial("tcp", "192.168.1.102:22", sshConfig)

	// create device using that connection which will be used by gomiko to start session
	device, err := gomiko.NewDeviceFromClient(conn, "juniper")

	if err != nil {
		log.Fatal(err)
	}
	// Open Session to device.
	if err := device.OpenSession(); err != nil {
		log.Fatal(err)
	}

	// execut commands
	output1, _ := device.SendCommand("show interfaces brief")

	// send a set of config commands
	commands := []string{"set routing-options static route 192.168.47.0/24 next-hop 172.16.1.2"}
	result2, _ := device.SendConfigSet(commands)

	fmt.Println(output1)
	fmt.Println(result2)
}
