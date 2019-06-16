package arista

import (
	"gomiko/pkg/connections"
	"gomiko/pkg/lib"
	"gomiko/pkg/lib/cisco"
	"gomiko/pkg/types"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {
	connection := connections.NewConnection(Host, Username, Password, "ssh")
	driver := lib.Driver{connection,"\n"}
	//Arista is exactly same as Cisco.
	base := cisco.CSCODevice{Password, DeviceType, "", driver, connection}

	return &EOSDevice{Host, Username, Password, base}

}
