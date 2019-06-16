package cisco

import (
	"gomiko/pkg/connections"
	"gomiko/pkg/lib"
	"gomiko/pkg/types"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {
	connection := connections.NewConnection(Host, Username, Password, "ssh")
	driver := lib.Driver{connection, "\n"}
	base := CSCODevice{Host, Password, DeviceType, "", driver, connection}
	switch DeviceType {
	case "cisco_asa":
		return &ASADevice{Host, Username, Password, base}
	case "cisco_ios":
		return &IOSDevice{Host, Username, Password, base}
	case "cisco_nxos":
		return &NXOSDevice{Host, Username, Password, base}
	case "cisco_iosxr":
		return &IOSXRDevice{Host, Username, Password, base}

	}

	return &CSCODevice{Host, Password, DeviceType, "", driver, connection}
}
