package cisco

import (
	"testing"
)

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

func TestASADevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode

	mockb := mockBase{}
	asaDevice := ASADevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	asaDevice.Connect()

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestASADevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	asaDevice := ASADevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	asaDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestASADevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	asaDevice := ASADevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	asaDevice.SendCommand("cmd")

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestASADevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	asaDevice := ASADevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	asaDevice.SendConfigSet(cmds)

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}
