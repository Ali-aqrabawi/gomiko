package arista

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

type EOSDevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *EOSDevice) Connect() error {
	return d.base.Connect()

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

func (d *EOSDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)

}

func (d *EOSDevice) SetTimeout(timeout uint8) {
	d.base.SetTimeout(timeout)
}
