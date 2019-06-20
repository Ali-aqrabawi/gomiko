package mikrotik

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
	"log"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {
	Username += "+ct200w" // disable paging and disable coloring
	devDriver := driver.NewDriver(Host, Username, Password, "\r", "ssh")

	switch DeviceType {
	case "mikrotik_routeros":
		return &MikroTikRouterOS{Host, Password, DeviceType, "", devDriver}
	default:
		log.Fatal("unsupported DeviceType: ", DeviceType)

	}
	return nil

}
