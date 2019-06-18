package cisco

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/utils"
	"strings"
)

type CSCODevice struct {
	Host       string
	Password   string
	DeviceType string
	Prompt     string
	Driver     driver.IDriver
}

func (d *CSCODevice) Connect() {
	d.Driver.Connect()
	d.Prompt = d.Driver.FindDevicePrompt("\r?(.*)[#>]", "#|>")
	utils.LogInfo(d.Host, "prompt found: "+d.Prompt)
	d.sessionPreparation()

}

func (d *CSCODevice) Disconnect() {
	d.Driver.Disconnect()
}

func (d *CSCODevice) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)
	if err != nil {
		utils.LogFatal(d.Host, "failed to send command: "+cmd, err)
	}

	return result, err

}

func (d *CSCODevice) SendConfigSet(cmds []string) (string, error) {

	results, _ := d.Driver.SendCommand("config term", d.Prompt)

	cmds = append(cmds, "end")

	out, err := d.Driver.SendCommandsSet(cmds, d.Prompt)
	results += out

	return results, err

}

func (d *CSCODevice) sessionPreparation() {
	utils.LogInfo(d.Host, "session preparation started...")

	out, err := d.Driver.SendCommand("enable", "Password:|"+d.Prompt)
	if strings.Contains(out, "Password:") {
		out, err = d.Driver.SendCommand(d.Password, d.Prompt)
	}

	if !strings.Contains(out, "#") {
		utils.LogFatal(d.Host, "failed to enter enable mode, output: "+out, nil)
	}

	cmd := getPagerDisableCmd(d.DeviceType)

	out, err = d.SendCommand(cmd)

	if err != nil {
		utils.LogFatal(d.Host, "failed to disable pagination", err)
	}

	utils.LogInfo(d.Host, "device output: "+out)
	utils.LogInfo(d.Host, "session preparation done!")

}
