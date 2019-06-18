package cisco

import (
	"gomiko/pkg/driver"
	"gomiko/pkg/types"
)

type IOSDevice struct {
	Host     string
	Username string
	Password string
	Driver driver.IDriver
	base     types.Device
}

func (d *IOSDevice) Connect() {
	d.base.Connect()

}

func (d *IOSDevice) Disconnect() {

	d.base.Disconnect()

}

func (d *IOSDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *IOSDevice) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)

}
