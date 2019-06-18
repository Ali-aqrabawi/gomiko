package cisco

import (
	"testing"
)

func TestIOSDevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode
	mockb := mockBase{}
	iosDevice := IOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	iosDevice.Connect()

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestIOSDevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	iosDevice := IOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	iosDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestIOSDevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	iosDevice := IOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	iosDevice.SendCommand("cmd")

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestIOSDevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	iosDevice := IOSDevice{"host", "username", "password", mockDriver{}, &mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	iosDevice.SendConfigSet(cmds)

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}
