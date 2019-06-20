package gomiko

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/arista"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/cisco"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/juniper"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/mikrotik"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
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
