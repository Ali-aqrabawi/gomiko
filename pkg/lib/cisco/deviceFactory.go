package cisco

import (
	"errors"
	"github.com/Ali-aqrabawi/gomiko/pkg/connections"
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.CiscoDevice, error) {

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
		}, nil

	case "cisco_ios":
		return &IOSDevice{
			Driver: devDriver,
			Prompt: "",
			base:   &base,
		}, nil
	case "cisco_nxos":
		return &NXOSDevice{
			Driver: devDriver,
			Prompt: "",
			base:   &base,
		}, nil
	case "cisco_iosxr":
		return &IOSXRDevice{
			Driver: devDriver,
			Prompt: "",
			base:   &base,
		}, nil

	default:
		return nil, errors.New("unsupported DeviceType: " + DeviceType)

	}

}
