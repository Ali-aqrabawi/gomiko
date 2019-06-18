package cisco

import (
	"gomiko/pkg/driver"
	"gomiko/pkg/types"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {

	devDriver := driver.NewDriver(Host, Username, Password, "\n", "ssh")
	base := CSCODevice{Host, Password, DeviceType, "", devDriver}
	switch DeviceType {
	case "cisco_asa":
		return &ASADevice{Host, Username, Password, devDriver,&base}
	case "cisco_ios":
		return &IOSDevice{Host, Username, Password, devDriver,&base}
	case "cisco_nxos":
		return &NXOSDevice{Host, Username, Password, devDriver,&base}
	case "cisco_iosxr":
		return &IOSXRDevice{Host, Username, Password, devDriver,&base}

	}

	return &CSCODevice{Host, Password, DeviceType, "", devDriver}
}
