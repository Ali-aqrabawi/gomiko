package arista

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/connections"
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/cisco"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection,  DeviceType string) types.CiscoDevice {
	devDriver := driver.NewDriver(connection, "\n")

	base := cisco.NewDevice(connection,  DeviceType)

	return &EOSDevice{
		Driver: devDriver,
		Prompt: "",
		base:   base,
	}

}
