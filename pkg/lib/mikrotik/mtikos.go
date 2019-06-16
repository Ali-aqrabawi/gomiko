package mikrotik

import (
	"errors"
	"gomiko/pkg/connections"
	"gomiko/pkg/lib"
	"gomiko/pkg/utils"
	"strings"
)

type MikroTikROS struct {
	Host       string
	Password   string
	DeviceType string
	Prompt     string
	Driver     lib.Driver
	Connection connections.Connection
}

func (d *MikroTikROS) Connect() {

	d.Connection.Connect()
	d.Prompt = d.Driver.FindDevicePrompt("\\[.*(@.*\\] >)", "] >")
	logger.Log(d.Host, "prompt found: "+d.Prompt)

}

func (d *MikroTikROS) Disconnect() {

	d.Connection.Disconnect()

}

func (d *MikroTikROS) SendCommand(cmd string) (string, error) {
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

func (d *MikroTikROS) SendConfigSet(cmds []string) (string, error) {
	logger.Log(d.Host, "sending config set: "+strings.Join(cmds, ", "))
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	results, err := d.Driver.SendCommandsSet(cmds, d.Prompt)

	return results, err

}

//func (d *MikroTikROS) sessionPreparation() {
//	logger.Log(d.Host, "session preparation started...")
//	d.Driver.ReadUntil(d.Prompt)
//	logger.Log(d.Host, "session preparation done!")
//
//}
