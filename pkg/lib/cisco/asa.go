package cisco

import "gomiko/pkg/types"

type ASADevice struct {
	Host     string
	Username string
	Password string
	base     types.Device
}

func (d *ASADevice) Connect() {
	d.base.Connect()

}

func (d *ASADevice) Disconnect() {

	d.base.Disconnect()

}

func (d *ASADevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *ASADevice) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)
}
