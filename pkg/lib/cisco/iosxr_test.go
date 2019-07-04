package cisco

import (
	"testing"
)

func TestIOSXRDevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode
	mockb := mockBase{}
	iosxrDevice := IOSXRDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	if err := iosxrDevice.Connect(); err != nil {
		t.Fatal(err)
	}

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestIOSXRDevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	iosxrDevice := IOSXRDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	iosxrDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestIOSXRDevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	iosxrDevice := IOSXRDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	_, err := iosxrDevice.SendCommand("cmd")
	if err != nil {
		panic(err)
	}

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestIOSXRDevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	iosxrDevice := IOSXRDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	_, err := iosxrDevice.SendConfigSet(cmds)
	if err != nil {
		panic(err)
	}

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}
