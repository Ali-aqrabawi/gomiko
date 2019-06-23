package cisco

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/connections"
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"log"
)

func NewDevice(connection connections.Connection,  DeviceType string) CiscoDevice {

	devDriver := driver.NewDriver(connection, "\n")
	base := CSCODevice{
		Driver:     devDriver,
		Prompt:     "",
		DeviceType: DeviceType,
	}
	switch DeviceType {
	case "cisco_asa":
		return &ASADevice{
			Driver: devDriver,
			Prompt: "",
			base:   &base,
		}
	case "cisco_ios":
		return &IOSDevice{
			Driver: devDriver,
			Prompt: "",
			base:   &base,}
	case "cisco_nxos":
		return &NXOSDevice{
			Driver: devDriver,
			Prompt: "",
			base:   &base,}
	case "cisco_iosxr":
		return &IOSXRDevice{
			Driver: devDriver,
			Prompt: "",
			base:   &base,}

	default:
		log.Fatal("unsupported DeviceType: ", DeviceType)

	}

	return nil
}
