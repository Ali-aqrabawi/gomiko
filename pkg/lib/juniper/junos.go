package juniper

import (
	"errors"
	"gomiko/pkg/connections"
	"gomiko/pkg/lib"
	"gomiko/pkg/utils"
	"strings"
)

type JunOSDevice struct {
	Host       string
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
	logger.Log(d.Host, "prompt found: "+d.Prompt)
	d.sessionPreparation()

}

func (d *JunOSDevice) Disconnect() {

	d.Connection.Disconnect()

}

func (d *JunOSDevice) SendCommand(cmd string) (string, error) {
	logger.Log(d.Host, "sending command: "+cmd)
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err

}

func (d *JunOSDevice) SendConfigSet(cmds []string) (string, error) {
	logger.Log(d.Host, "sending config set: "+strings.Join(cmds, ", "))
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
	logger.Log(d.Host, "session preparation started...")
	out, err := d.Driver.SendCommand("cli", d.Prompt)
	if !strings.Contains(out, ">") {
		logger.Fatal(d.Host, "failed to enter enable mode", nil)
	}

	out, err = d.SendCommand("set cli screen-length 0")

	if err != nil {
		logger.Fatal(d.Host, "failed to disable pagination", err)
	}
	logger.Log(d.Host, "session preparation done!")
	logger.Log(d.Host, "device output: "+out)

}
