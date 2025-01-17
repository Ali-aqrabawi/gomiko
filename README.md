# Gomiko
[![Build Status](https://travis-ci.org/Ali-aqrabawi/gomiko.svg?branch=master)](https://travis-ci.org/Ali-aqrabawi/gomiko)
[![GolangCI](https://golangci.com/badges/github.com/Ali-aqrabawi/gomiko.svg)](https://golangci.com)
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/Ali-aqrabawi/gomiko)
[<img src="https://api.gitsponsors.com/api/badge/img?id=192590075" height="20">](https://api.gitsponsors.com/api/badge/link?p=xOtJljV6r7vHqTB7WwqJPvHjy9mv5zm1vUkLp13Uar0JcZYpgrTgWJHMoakqNuul)




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
* Nokia SROS

## Installation
get gomiko pkg: `go get -u github.com/Ali-aqrabawi/gomiko/pkg`.

## Examples :
```go
import (
	"fmt"
	"log"
	"github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	
     device, err := gomiko.NewDevice("192.168.1.1", "admin", "password", "cisco_ios", 22)
     
     if err != nil {
     	log.Fatal(err)
     }
     
     //Connect to device
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
```

## create device with enable password:
```go
import (
	"fmt"
	"log"
	"github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	
     device, err := gomiko.NewDevice("192.168.1.1", "admin", "password", "cisco_ios", 22, gomiko.SecretOption("enablePass"))
     
     if err != nil {
     	log.Fatal(err)
     }     

}
```

## create device with custom timeout:
```go
import (
	"fmt"
	"log"
	"github.com/Ali-aqrabawi/gomiko/pkg"
)

func main() {
	
     device, err := gomiko.NewDevice("192.168.1.1", "admin", "password", "cisco_ios", 22, gomiko.SecretOption("enablePass"), gomiko.TimeoutOption(10))
     
     if err != nil {
     	log.Fatal(err)
     }     

}
```
