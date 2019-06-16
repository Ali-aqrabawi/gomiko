package gomiko

import (
	"gomiko/pkg/lib/arista"
	"gomiko/pkg/lib/cisco"
	"gomiko/pkg/lib/juniper"
	"gomiko/pkg/lib/mikrotik"
	"gomiko/pkg/types"
	"strings"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {
	if strings.Contains(DeviceType, "cisco") {
		return cisco.NewDevice(Host, Username, Password, DeviceType)
	} else if strings.Contains(DeviceType, "arista") {
		return arista.NewDevice(Host, Username, Password, DeviceType)
	} else if strings.Contains(DeviceType, "juniper") {
		return juniper.NewDevice(Host, Username, Password, DeviceType)
	} else if strings.Contains(DeviceType, "mikrotik") {
		return mikrotik.NewDevice(Host, Username, Password, DeviceType)
	} else {
		panic("DeviceType not supported: " + DeviceType)
	}

}
