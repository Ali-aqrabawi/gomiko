package mikrotik

import (
	"errors"
	"github.com/Ali-aqrabawi/gomiko/pkg/connections"
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.Device, error) {
	devDriver := driver.NewDriver(connection, "\r")

	switch DeviceType {
	case "mikrotik_routeros":
		return &MikroTikRouterOS{
			Driver:     devDriver,
			DeviceType: DeviceType,
			Prompt:     "",
		}, nil
	default:
		return nil, errors.New("unsupported DeviceType: " + DeviceType)

	}

}
