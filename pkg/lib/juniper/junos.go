package juniper

import (
	"errors"
	"gomiko/pkg/connections"
	"gomiko/pkg/lib"
	"strings"
)

type JunOSDevice struct {
	Password   string
	DeviceType string
	Prompt     string
	Driver     lib.Driver
	Connection connections.Connection
}

func (d *JunOSDevice) Connect() {

	d.Connection.Connect()
	d.Driver.ReadUntil("% ")
	d.Prompt = d.Driver.FindDevicePrompt("(@.*)[#>%]", "%")

	d.sessionPreparation()

}

func (d *JunOSDevice) Disconnect() {

	d.Connection.Disconnect()

}

func (d *JunOSDevice) SendCommand(cmd string) (string, error) {
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err

}

func (d *JunOSDevice) SendConfigSet(cmds []string) (string, error) {
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	results, _ := d.Driver.SendCommand("configure", d.Prompt)

	cmds = append(cmds, "commit", "exit")

	out, err := d.Driver.SendCommandsSet(cmds, d.Prompt)
	results += out
	return results, err

}

func (d *JunOSDevice) sessionPreparation() {

	out, err := d.Driver.SendCommand("cli", d.Prompt)
	if !strings.Contains(out, ">") {
		panic("failed to enter cli mode: " + out)
	}

	out, err = d.SendCommand("set cli screen-length 0")

	if err != nil {
		panic(err)
	}

}
