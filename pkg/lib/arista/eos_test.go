package arista

import (
	"testing"
)

type mockBase struct {
	Calls *string
}

func (b mockBase) Connect() error {
	*b.Calls = "Connect"
	return nil

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
func (b mockBase) SetSecret(secret string) {
	*b.Calls = "SetSecret"

}

func (b mockBase) SetTimeout(timeout uint8) {
}

type mockDriver struct {
	ReadSideEffect func() string
	CmdCalls       *string
	PatternCalls   *string
	PromptRegex    *string
	GenericCalls   *string
}

func (c mockDriver) Connect() error {
	return nil

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

	return c.ReadUntil(expectPattern)

}
func (c mockDriver) SendCommandsSet(cmds []string, expectPattern string) (string, error) {
	for _, cmd := range cmds {
		_, err := c.SendCommand(cmd, expectPattern)
		if err != nil {
			panic(err)
		}
	}
	return c.ReadUntil(expectPattern)

}

func (c mockDriver) FindDevicePrompt(regex string, pattern string) (string, error) {
	*c.PromptRegex = regex
	return c.ReadUntil(pattern)

}

func (c mockDriver) ReadUntil(pattern string) (string, error) {
	*c.PatternCalls += pattern + ", "

	return c.ReadSideEffect(), nil

}

func (c mockDriver) SetTimeout(timeout uint8) {
}

func TestEOSDevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode

	mockb := mockBase{}
	eosDevice := EOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	if err := eosDevice.Connect(); err != nil {
		t.Fatal(err)
	}

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestEOSDevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	eosDevice := EOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	eosDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestEOSDevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	eosDevice := EOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	_, err := eosDevice.SendCommand("cmd")
	if err != nil {
		panic(err)
	}

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestEOSDevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	eosDevice := EOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	_, err := eosDevice.SendConfigSet(cmds)
	if err != nil {
		panic(err)
	}

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}
