package mikrotik

import (
	"errors"
	"gomiko/pkg/connections"
	"gomiko/pkg/lib"
)

type MikroTikROS struct {
	Password   string
	DeviceType string
	Prompt     string
	Driver     lib.Driver
	Connection connections.Connection
}

func (d *MikroTikROS) Connect() {

	d.Connection.Connect()
	d.Prompt = d.Driver.FindDevicePrompt("\\[.*(@.*\\] >)", "] >")

}

func (d *MikroTikROS) Disconnect() {

	d.Connection.Disconnect()

}

func (d *MikroTikROS) SendCommand(cmd string) (string, error) {
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err

}

func (d *MikroTikROS) SendConfigSet(cmds []string) (string, error) {
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	results, err := d.Driver.SendCommandsSet(cmds, d.Prompt)

	return results, err

}

func (d *MikroTikROS) sessionPreparation() {

	d.Driver.ReadUntil(d.Prompt)

}
