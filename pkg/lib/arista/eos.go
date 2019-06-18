package arista

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

type EOSDevice struct {
	Host     string
	Username string
	Password string
	Driver driver.IDriver
	base     types.Device
}

func (d *EOSDevice) Connect() {
	d.base.Connect()

}

func (d *EOSDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *EOSDevice) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)

}

func (d *EOSDevice) Disconnect() {
	d.base.Disconnect()

}
