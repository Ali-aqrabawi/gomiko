package cisco

import (
	"errors"
	"gomiko/pkg/connections"
	"gomiko/pkg/lib"
	"gomiko/pkg/utils"
	"strings"
)

type CSCODevice struct {
	Host       string
	Password   string
	DeviceType string
	Prompt     string
	Driver     lib.Driver
	Connection connections.Connection
}

func (d *CSCODevice) Connect() {

	d.Connection.Connect()
	d.Prompt = d.Driver.FindDevicePrompt("\r?(.*)[#>]", "#|>")
	logger.Log(d.Host, "prompt found: " + d.Prompt)
	d.sessionPreparation()

}

func (d *CSCODevice) Disconnect() {
	d.Connection.Disconnect()

}

func (d *CSCODevice) SendCommand(cmd string) (string, error) {
	logger.Log(d.Host, "sending command: "+cmd)
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	result, err := d.Driver.SendCommand(cmd, d.Prompt)
	if err != nil {
		logger.Fatal(d.Host, "failed to send command: "+cmd, err)
	}

	return result, err

}

func (d *CSCODevice) SendConfigSet(cmds []string) (string, error) {
	logger.Log(d.Host, "sending config set: "+strings.Join(cmds, ", "))
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
	logger.Log(d.Host, "session preparation started...")

	out, err := d.Driver.SendCommand("enable", "Password:|"+d.Prompt)
	out, err = d.Driver.SendCommand(d.Password, d.Prompt)

	if !strings.Contains(out, "#") {
		logger.Fatal(d.Host, "failed to enter enable mode", nil)
	}

	cmd := getPagerDisableCmd(d.DeviceType)

	out, err = d.SendCommand(cmd)

	if err != nil {
		logger.Fatal(d.Host, "failed to disable pagination", err)
	}

	logger.Log(d.Host, "device output: "+out)
	logger.Log(d.Host, "session preparation done!")

}
