gomiko
******

multi-vendor networking SDK.

Supports:
---------
* Cisco IOS
* Cisco IOS XR
* Cisco ASA
* Cisco NX-OS
* Mikrotik RouterOS
* Arista EOS
* Juniper JunOS

Features:
---------
* SSH
* Telnet
* TextFSM

Examples:
---------
Example1 :

.. code-block:: go

    import (
    	"fmt"
    	"gomiko/pkg"
    )

    func main() {

    	device := gomiko.NewDevice("192.168.1.99", "admin", "pass", "cisco_asa")

    	device.Connect()


    	cmds := []string{"object network GoLangObj","host 44.6.3.1"}

    	result, _ := device.SendConfigSet(cmds)

    	fmt.Println(result)


