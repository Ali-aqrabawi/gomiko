package cisco

import (
	"errors"
	"gomiko/pkg/connections"
	"gomiko/pkg/lib"
	"strings"
)

type CSCODevice struct {
	Password   string
	DeviceType string
	Prompt     string
	Driver     lib.Driver
	Connection connections.Connection
}

func (d *CSCODevice) Connect() {

	d.Connection.Connect()
	d.Prompt = d.Driver.FindDevicePrompt("(.*)[#>]", "")
	d.sessionPreparation()

}

func (d *CSCODevice) Disconnect() {

	d.Connection.Disconnect()

}

func (d *CSCODevice) SendCommand(cmd string) (string, error) {
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err

}

func (d *CSCODevice) SendConfigSet(cmds []string) (string, error) {
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	results, _ := d.Driver.SendCommand("config term", d.Prompt)

	cmds = append(cmds, "end")

	out, err := d.Driver.SendCommandsSet(cmds, d.Prompt)
	results += out

	return results, err

}

func (d *CSCODevice) sessionPreparation() {

	out, err := d.Driver.SendCommand("enable", "Password:")
	out, err = d.Driver.SendCommand(d.Password, d.Prompt)

	if !strings.Contains(out, "#") {
		panic("failed to enter enable mode: " + out)
	}

	cmd := getPagerDisableCmd(d.DeviceType)

	out, err = d.SendCommand(cmd)

	if err != nil {
		panic(err)
	}

}
