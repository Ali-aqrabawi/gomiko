package arista

import "testing"

type mockBase struct {
	Calls *string
}

func (b mockBase) Connect() {
	*b.Calls = "Connect"

}
func (b mockBase) Disconnect() {
	*b.Calls = "Disconnect"

}
func (b mockBase) SendCommand(cmd string) (string, error) {
	*b.Calls = "SendCommand"

	return "", nil

}
func (b mockBase) SendConfigSet(cmds []string) (string, error) {
	*b.Calls = "SendConfigSet"
	return "", nil

}

func (b mockBase) sessionPreparation() {
	*b.Calls = "sessionPreparation"

}

type mockDriver struct {
	ReadSideEffect func() string
	CmdCalls       *string
	PatternCalls   *string
	PromptRegex    *string
	GenericCalls   *string
}

func (c mockDriver) Connect() {

}
func (c mockDriver) Disconnect() {
	*c.GenericCalls = "disconnect"

}
func (c mockDriver) SendCommand(cmd string, expectPattern string) (string, error) {
	if c.CmdCalls != nil {
		*c.CmdCalls += cmd + ", "
	} else {
		c.CmdCalls = &cmd
	}

	return c.ReadUntil(expectPattern), nil

}
func (c mockDriver) SendCommandsSet(cmds []string, expectPattern string) (string, error) {
	for _, cmd := range cmds {
		c.SendCommand(cmd, expectPattern)
	}
	return c.ReadUntil(expectPattern), nil

}

func (c mockDriver) FindDevicePrompt(regex string, pattern string) string {
	*c.PromptRegex = regex
	return c.ReadUntil(pattern)

}

func (c mockDriver) ReadUntil(pattern string) string {
	*c.PatternCalls += pattern + ", "

	return c.ReadSideEffect()

}

func TestASADevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode

	mockb := mockBase{}
	asaDevice := EOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	asaDevice.Connect()

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestASADevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	asaDevice := EOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	asaDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestASADevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	asaDevice := EOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	asaDevice.SendCommand("cmd")

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestASADevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	asaDevice := EOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	asaDevice.SendConfigSet(cmds)

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}
