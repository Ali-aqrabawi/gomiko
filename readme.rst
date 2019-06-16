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
Example :

.. code-block:: go

    import (
    	"fmt"
    	"gomiko/pkg"
    )

    func main() {

    	device := gomiko.NewDevice("162.221.4.102", "admin", "J3llyfish1", "cisco_asa")
    	fmt.Println("werw")

    	device.Connect()


    	cmds := []string{"object network GoLangObj","host 44.6.3.1"}

    	result, _ := device.SendConfigSet(cmds)

    	fmt.Println(result)


