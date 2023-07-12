package cisco

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

func TestASADevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode

	mockb := mockBase{}
	asaDevice := ASADevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	if err := asaDevice.Connect(); err != nil {
		t.Fatal(err)
	}

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestASADevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	asaDevice := ASADevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	asaDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestASADevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	asaDevice := ASADevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	_, err := asaDevice.SendCommand("cmd")
	if err != nil {
		panic(err)
	}

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestASADevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	asaDevice := ASADevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	_, err := asaDevice.SendConfigSet(cmds)
	if err != nil {
		panic(err)
	}

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}
