package cisco

import (
	"testing"
)

func TestNXOSDevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode
	mockb := mockBase{}
	nxosDevice := NXOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	nxosDevice.Connect()

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestNXOSDevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	nxosDevice := NXOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	nxosDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestNXOSDevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	nxosDevice := NXOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	nxosDevice.SendCommand("cmd")

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestNXOSDevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	nxosDevice := NXOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	nxosDevice.SendConfigSet(cmds)

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}
