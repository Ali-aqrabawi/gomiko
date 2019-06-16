package mikrotik

import (
	"gomiko/pkg/connections"
	"gomiko/pkg/lib"
	"gomiko/pkg/types"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {
	Username += "+ct200w" // disable paging and disable coloring
	connection := connections.NewConnection(Host, Username, Password, "ssh")
	driver := lib.Driver{connection, "\r"}
	return &MikroTikROS{Host, Password, DeviceType, "", driver, connection}

}
