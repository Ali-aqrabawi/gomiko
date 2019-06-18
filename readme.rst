gomiko
******

multi-vendor networking SDK.
inspired by `netmiko`

Supports:
---------
* Cisco IOS
* Cisco IOS XR
* Cisco ASA
* Cisco NX-OS
* Mikrotik RouterOS
* Arista EOS
* Juniper JunOS

Quick Start:
------------
get gomiko pkg:

.. code:: bash

    go get -u github.com/Ali-aqrabawi/gomiko/pkg


Examples:
---------
Example1 sending command:

.. code-block:: go

    import (
    	"fmt"
        "github.com/Ali-aqrabawi/gomiko/pkg"
    )

    func main() {

    	device := gomiko.NewDevice("192.168.1.99", "admin", "pass", "cisco_asa")

    	device.Connect()


    	}

    	result, _ := device.SendCommand("show version")

        device.Disconnect()

    	fmt.Println(result)


Example2 sending config list:

.. code-block:: go

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


