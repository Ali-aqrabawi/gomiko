package cisco

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
)

type IOSDevice struct {
	Driver driver.IDriver
	Prompt string
	base   CiscoDevice
}

func (d *IOSDevice) Connect() error {
	return d.base.Connect()

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
func (d *IOSDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)
}
