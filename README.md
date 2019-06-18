# Gomiko
[![Build Status](https://travis-ci.org/Ali-aqrabawi/gomiko.svg?branch=master)](https://travis-ci.org/Ali-aqrabawi/gomiko)
[![GolangCI](https://golangci.com/badges/github.com/Ali-aqrabawi/gomiko.svg)](https://golangci.com)

multi-vendor networking SDK.
inspired by `netmiko`.


## Quick Start

get gomiko pkg `go get -u github.com/Ali-aqrabawi/gomiko/pkg`

 
## supports
* Cisco IOS
* Cisco IOS XR
* Cisco ASA
* Cisco NX-OS
* Mikrotik RouterOS
* Arista EOS
* Juniper JunOS

## Examples
 
 1. execute a command:

        import (
    	    "fmt"
            "github.com/Ali-aqrabawi/gomiko/pkg"
         )

        func main() {

    	    device := gomiko.NewDevice("192.168.1.99", "admin", "pass", "cisco_asa")

    	    device.Connect()
    	    	
    	    result, _ := device.SendCommand("show version")

            device.Disconnect()

    	    fmt.Println(result)


    	    }

    

 2. execute config set:

        import (
    	    "fmt"
    	    "github.com/Ali-aqrabawi/gomiko/pkg"
        )

        func main() {

    	    device := gomiko.NewDevice("192.168.1.99", "admin", "pass", "cisco_asa")

    	    device.Connect()


    	    cmds := []string{"object network GoLangObj","host 44.6.3.1"}

    	    result, _ := device.SendConfigSet(cmds)

            device.Disconnect()

    	    fmt.Println(result)
    	}
