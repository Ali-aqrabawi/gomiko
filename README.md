# Gomiko
[![Build Status](https://travis-ci.org/Ali-aqrabawi/gomiko.svg?branch=master)](https://travis-ci.org/Ali-aqrabawi/gomiko)
[![GolangCI](https://golangci.com/badges/github.com/Ali-aqrabawi/gomiko.svg)](https://golangci.com)
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/Ali-aqrabawi/gomiko)

Gomiko is a `Go` implementation of [netmiko](https://github.com/ktbyers/netmiko). It serves as multi-vendor networking SDK that helps communicate and execute commands via an interactive `shell`
without needing to care about handling device prompts and terminal modes.
 
## Supports
* Cisco IOS
* Cisco IOS XR
* Cisco ASA
* Cisco NX-OS
* Mikrotik RouterOS
* Arista EOS
* Juniper JunOS

## Installation
get gomiko pkg: `go get -u github.com/Ali-aqrabawi/gomiko/pkg`.

## Examples 
 1. create device using basic parameters and execute commands:
```go
import (
	"fmt"
	"github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	
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
```

 2. create device using `*ssh.Clinet` and execute commands:
```go
import (
	"fmt"
	"github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	
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
```

