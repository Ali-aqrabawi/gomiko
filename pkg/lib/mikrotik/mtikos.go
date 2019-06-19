package mikrotik

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
)

type MikroTikROS struct {
	Host       string
	Password   string
	DeviceType string
	Prompt     string
	Driver     driver.IDriver
}

func (d *MikroTikROS) Connect() error {

	if err := d.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := d.Driver.FindDevicePrompt("\\[.*(@.*\\] >)", "] >")
	if err != nil {
		return err
	}
	d.Prompt = prompt
	return d.sessionPreparation()

}

func (d *MikroTikROS) Disconnect() {

	d.Driver.Disconnect()

}

func (d *MikroTikROS) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)


	return result, err

}

func (d *MikroTikROS) SendConfigSet(cmds []string) (string, error) {

	results, err := d.Driver.SendCommandsSet(cmds, d.Prompt)

	return results, err

}

func (d *MikroTikROS) sessionPreparation() error {
	_, err := d.Driver.SendCommand("", d.Prompt)
	return err

}
