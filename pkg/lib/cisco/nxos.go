package cisco

import (
	"gomiko/pkg/driver"
	"gomiko/pkg/types"
)

type NXOSDevice struct {
	Host     string
	Username string
	Password string
	Driver   driver.IDriver
	base     types.Device
}

func (d *NXOSDevice) Connect() {
	d.base.Connect()

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
