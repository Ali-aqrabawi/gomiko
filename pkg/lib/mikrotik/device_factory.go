package mikrotik

import (
	"gomiko/pkg/driver"
	"gomiko/pkg/types"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {
	Username += "+ct200w" // disable paging and disable coloring
	devDriver := driver.NewDriver(Host, Username, Password, "\r", "ssh")

	return &MikroTikROS{Host, Password, DeviceType, "", devDriver}

}
