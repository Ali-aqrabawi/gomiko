package arista

import "gomiko/pkg/lib/cisco"

type EOSDevice struct {
	Host     string
	Username string
	Password string
	base     cisco.CSCODevice
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