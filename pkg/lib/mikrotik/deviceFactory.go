package mikrotik

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/connections"
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
	"log"
)

func NewDevice(connection connections.Connection, DeviceType string) types.Device {
	devDriver := driver.NewDriver(connection, "\r")

	switch DeviceType {
	case "mikrotik_routeros":
		return &MikroTikRouterOS{
			Driver:     devDriver,
			DeviceType: DeviceType,
			Prompt:     "",
		}
	default:
		log.Fatal("unsupported DeviceType: ", DeviceType)

	}
	return nil

}
