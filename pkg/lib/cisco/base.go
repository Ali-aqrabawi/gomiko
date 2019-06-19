package cisco

import (
	"errors"
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"strings"
)

type CSCODevice struct {
	Host       string
	Password   string
	DeviceType string
	Prompt     string
	Driver     driver.IDriver
}

func (d *CSCODevice) Connect() error {
	if err := d.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := d.Driver.FindDevicePrompt("\r?(.*)[#>]", "#|>")

	if err != nil {
		return err
	}
	d.Prompt = prompt
	return d.sessionPreparation()

}

func (d *CSCODevice) Disconnect() {
	d.Driver.Disconnect()
}

func (d *CSCODevice) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)


	return result, err

}

func (d *CSCODevice) SendConfigSet(cmds []string) (string, error) {

	results, _ := d.Driver.SendCommand("config term", d.Prompt)

	cmds = append(cmds, "end")

	out, err := d.Driver.SendCommandsSet(cmds, d.Prompt)
	results += out

	return results, err

}

func (d *CSCODevice) sessionPreparation() error {

	out, err := d.Driver.SendCommand("enable", "Password:|"+d.Prompt)
	if err != nil {
		return errors.New("failed to send enable command:" + err.Error())
	}
	if strings.Contains(out, "Password:") {
		out, err = d.Driver.SendCommand(d.Password, d.Prompt)
		if err != nil {
			return errors.New("failed to send enable password:" + err.Error())
		}
	}

	if !strings.Contains(out, "#") {
		return errors.New("failed to enter enable mode")
	}

	cmd := getPagerDisableCmd(d.DeviceType)

	out, err = d.SendCommand(cmd)

	return err

}
