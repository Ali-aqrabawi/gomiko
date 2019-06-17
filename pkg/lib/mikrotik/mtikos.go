package mikrotik

import (
	"gomiko/pkg/driver"
	"gomiko/pkg/utils"
)

type MikroTikROS struct {
	Host       string
	Password   string
	DeviceType string
	Prompt     string
	Driver     driver.IDriver
}

func (d *MikroTikROS) Connect() {

	d.Driver.Connect()
	d.Prompt = d.Driver.FindDevicePrompt("\\[.*(@.*\\] >)", "] >")
	logger.Log(d.Host, "prompt found: "+d.Prompt)

}

func (d *MikroTikROS) Disconnect() {

	d.Driver.Disconnect()

}

func (d *MikroTikROS) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)
	if err != nil {
		logger.Fatal(d.Host, "failed to send command: "+cmd, err)
	}

	return result, err

}

func (d *MikroTikROS) SendConfigSet(cmds []string) (string, error) {

	results, err := d.Driver.SendCommandsSet(cmds, d.Prompt)

	return results, err

}
