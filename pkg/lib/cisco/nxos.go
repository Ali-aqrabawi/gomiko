package cisco

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

type NXOSDevice struct {
	Host     string
	Username string
	Password string
	Driver   driver.IDriver
	base     types.Device
}

func (d *NXOSDevice) Connect() error{
	return d.base.Connect()

}

func (d *NXOSDevice) Disconnect() {
	d.base.Disconnect()

}

func (d *NXOSDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *NXOSDevice) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)

}
