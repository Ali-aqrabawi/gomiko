package juniper

import (
	"errors"
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"strings"
)

type JunOSDevice struct {
	Driver     driver.IDriver
	DeviceType string
	Prompt     string


}

func (d *JunOSDevice) Connect() error {
	if err := d.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := d.Driver.FindDevicePrompt("(@.*)[#>%]", "%")
	if err != nil {
		return err
	}
	d.Prompt = prompt

	return d.sessionPreparation()

}

func (d *JunOSDevice) Disconnect() {

	d.Driver.Disconnect()

}

func (d *JunOSDevice) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err

}

func (d *JunOSDevice) SendConfigSet(cmds []string) (string, error) {

	results, _ := d.Driver.SendCommand("configure", d.Prompt)

	cmds = append(cmds, "commit", "exit")

	out, err := d.Driver.SendCommandsSet(cmds, d.Prompt)
	results += out
	return results, err

}

func (d *JunOSDevice) sessionPreparation() error {

	out, err := d.Driver.SendCommand("cli", d.Prompt)
	if err != nil {
		return errors.New("failed to send cli command" + err.Error())
	}
	if !strings.Contains(out, ">") {
		return errors.New("failed to enter cli mode, device output: " + out)
	}

	out, err = d.SendCommand("set cli screen-length 0")

	if err != nil {
		return errors.New("failed to disable pagination" + err.Error())
	}
	return nil

}
