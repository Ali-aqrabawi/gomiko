package arista

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/cisco"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

func NewDevice(Host string, Username string, Password string, DeviceType string) types.Device {
	devDriver := driver.NewDriver(Host, Username, Password, "\n", "ssh")
	//Arista is exactly same as Cisco.
	base := cisco.CSCODevice{Host, Password, DeviceType, "", devDriver}

	return &EOSDevice{Host, Username, Password, devDriver,&base}

}
