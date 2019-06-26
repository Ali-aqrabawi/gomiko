package cisco

import (
	"testing"
)

type mockBase struct {
	Calls *string
}

func (b mockBase) OpenSession() error {
	*b.Calls = "OpenSession"
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

func TestASADevice_OpenSession(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode

	mockb := mockBase{}
	asaDevice := ASADevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	if err := asaDevice.OpenSession(); err != nil {
		t.Fatal(err)
	}

	if calls != "OpenSession" {
		t.Error("base.OpenSession() was not called")
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
