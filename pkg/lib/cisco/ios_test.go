package cisco

import (
	"testing"
)

func TestIOSDevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode
	mockb := mockBase{}
	iosDevice := IOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	if err := iosDevice.Connect(); err != nil {
		t.Fatal(err)
	}

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestIOSDevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	iosDevice := IOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	iosDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestIOSDevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	iosDevice := IOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	_, err := iosDevice.SendCommand("cmd")
	if err != nil {
		panic(err)
	}

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestIOSDevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	iosDevice := IOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	_, err := iosDevice.SendConfigSet(cmds)
	if err != nil {
		panic(err)
	}

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}
