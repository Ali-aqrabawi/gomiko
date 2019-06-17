package arista

import (
	"gomiko/pkg/driver"
	"gomiko/pkg/lib/cisco"
	"gomiko/pkg/types"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {
	devDriver := driver.NewDriver(Host, Username, Password, "\n", "ssh")
	//Arista is exactly same as Cisco.
	base := cisco.CSCODevice{Host, Password, DeviceType, "", devDriver}

	return &EOSDevice{Host, Username, Password, base}

}
