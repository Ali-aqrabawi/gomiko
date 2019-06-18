package juniper

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/utils"
	"strings"
)

type JunOSDevice struct {
	Host       string
	Password   string
	DeviceType string
	Prompt     string
	Driver     driver.IDriver
}

func (d *JunOSDevice) Connect() {
	d.Driver.Connect()
	d.Prompt = d.Driver.FindDevicePrompt("(@.*)[#>%]", "%")
	utils.LogInfo(d.Host, "prompt found: "+d.Prompt)
	d.sessionPreparation()

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

func (d *JunOSDevice) sessionPreparation() {
	utils.LogInfo(d.Host, "session preparation started...")

	out, err := d.Driver.SendCommand("cli", d.Prompt)
	if err != nil {
		utils.LogFatal(d.Host, "failed to send cli command", err)
	}
	if !strings.Contains(out, ">") {
		utils.LogFatal(d.Host, "failed to enter cli mode", nil)
	}

	out, err = d.SendCommand("set cli screen-length 0")

	if err != nil {
		utils.LogFatal(d.Host, "failed to disable pagination", err)
	}

	utils.LogInfo(d.Host, "device output: "+out)
	utils.LogInfo(d.Host, "session preparation done!")

}
