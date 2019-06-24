package gomiko

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/connections"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/arista"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/cisco"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/juniper"
	"github.com/Ali-aqrabawi/gomiko/pkg/lib/mikrotik"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"strings"
)

func NewDevice(Host string, Username string, Password string, DeviceType string, Port uint8, Options ...DeviceOption) (types.Device, error) {
	var device types.Device

	//for Mikrotik you need to append +ct200w to username
	if strings.Contains(DeviceType, "mikrotik") {
		Username += "+ct200w"
	}

	//create connection
	connection, err := connections.NewSSHConn(Host, Username, Password, Port)
	if err != nil {
		return nil, err
	}

	//create the Device
	if strings.Contains(DeviceType, "cisco") {
		device = cisco.NewDevice(&connection, DeviceType)
	} else if strings.Contains(DeviceType, "arista") {
		device = arista.NewDevice(&connection, DeviceType)
	} else if strings.Contains(DeviceType, "juniper") {
		device = juniper.NewDevice(&connection, DeviceType)
	} else if strings.Contains(DeviceType, "mikrotik") {
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

func NewDeviceFromClient(client *ssh.Client, DeviceType string, Options ...DeviceOption) (types.Device, error) {
	var device types.Device

	connection, err := connections.NewConnectionFromClient(client)
	if err != nil {
		return nil, err
	}
	if strings.Contains(DeviceType, "cisco") {
		device = cisco.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "arista") {
		device = arista.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "juniper") {
		device = juniper.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "mikrotik") {
		device = mikrotik.NewDevice(connection, DeviceType)
	} else {
		return nil, errors.New("DeviceType not supported: " + DeviceType)
	}
	for _, option := range Options {
		err := option(device)
		if err != nil {
			return nil, err
		}
	}
	return device, nil
}
