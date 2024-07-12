package arista

import (
	"errors"

	"github.com/asadarafat/gomiko/pkg/connections"
	"github.com/asadarafat/gomiko/pkg/driver"
	"github.com/asadarafat/gomiko/pkg/lib/cisco"
	"github.com/asadarafat/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.CiscoDevice, error) {
	devDriver := driver.NewDriver(connection, "\n")

	base := cisco.CSCODevice{
		Driver:     devDriver,
		DeviceType: "cisco_ios",
	}
	if DeviceType != "arista_eos" {
		return nil, errors.New("unsupported Arista device type: " + DeviceType)

	}

	return &EOSDevice{
		Driver: devDriver,
		base:   &base,
	}, nil

}
