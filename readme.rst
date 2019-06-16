gomiko
******
multi-vendor network SDK.

Requires:
---------
* asyncio
* AsyncSSH
* Python >=3.5
* pyYAML
* asyncssh

Supports:
---------
* Cisco IOS
* Cisco IOS XR
* Cisco ASA
* Cisco NX-OS


Features:
---------
* SSH
* Telnet
* TextFSM

Examples:
---------
Example of interacting with Cisco IOS devices:

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


