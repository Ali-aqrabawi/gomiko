package juniper

import (
	"gomiko/pkg/connections"
	"gomiko/pkg/lib"
	"gomiko/pkg/types"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {
	connection := connections.NewConnection(Host, Username, Password, "ssh")
	driver := lib.Driver{connection, "\n"}
	return &JunOSDevice{Host, Password, DeviceType, "", driver, connection}

}
