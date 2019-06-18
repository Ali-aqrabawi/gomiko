package cisco

import (
	"testing"
)



func TestIOSXRDevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode
	mockb := mockBase{}
	iosxrDevice := IOSXRDevice{"host", "username", "password", mockDriver{},&mockb}

	var calls string
	mockb.Calls = &calls
	iosxrDevice.Connect()

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestIOSXRDevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	iosxrDevice := IOSXRDevice{"host", "username", "password", mockDriver{},&mockb}

	var calls string
	mockb.Calls = &calls
	iosxrDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestIOSXRDevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	iosxrDevice := IOSXRDevice{"host", "username", "password", mockDriver{},&mockb}

	var calls string
	mockb.Calls = &calls
	iosxrDevice.SendCommand("cmd")

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestIOSXRDevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	iosxrDevice := IOSXRDevice{"host", "username", "password", mockDriver{},&mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	iosxrDevice.SendConfigSet(cmds)

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}
