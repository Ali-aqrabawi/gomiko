package juniper

import (
	"gomiko/pkg/driver"
	"gomiko/pkg/types"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {
	devDriver := driver.NewDriver(Host, Username, Password, "\n", "ssh")

	return &JunOSDevice{Host, Password, DeviceType, "", devDriver}

}
