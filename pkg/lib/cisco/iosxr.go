package cisco

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

type IOSXRDevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *IOSXRDevice) Connect() error {
	return d.base.Connect()

}

func (d *IOSXRDevice) Disconnect() {

	d.base.Disconnect()

}

func (d *IOSXRDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *IOSXRDevice) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)

}
func (d *IOSXRDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)
}
