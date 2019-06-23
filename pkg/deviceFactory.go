package gomiko

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/connections"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/arista"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/cisco"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/juniper"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/mikrotik"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
	"github.com/pkg/errors"
	"strings"
)

func NewDevice(Host string, Username string, Password string, DeviceType string, Options ...OptionDevice) (types.Device, error) {
	var device types.Device
	connection := connections.SSHConn{
		Host:     Host,
		Username: Username,
		Password: Password,
	}
	if strings.Contains(DeviceType, "cisco") {
		device = cisco.NewDevice(&connection, DeviceType)
	} else if strings.Contains(DeviceType, "arista") {
		device = arista.NewDevice(&connection, DeviceType)
	} else if strings.Contains(DeviceType, "juniper") {
		device = juniper.NewDevice(&connection, DeviceType)
	} else if strings.Contains(DeviceType, "mikrotik") {
		connection.Host += "+ct200w"
		device = mikrotik.NewDevice(&connection, DeviceType)
	} else {
		return nil, errors.New("DeviceType not supported: " + DeviceType)
	}

	// running Options Functions.
	for _, option := range Options {
		err := option(device)
		if err != nil {
			return nil, err
		}
	}

	return device, nil

}
