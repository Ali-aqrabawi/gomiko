package cisco

import (
	"testing"
)

func TestNXOSDevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode
	mockb := mockBase{}
	nxosDevice := NXOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	if err :=nxosDevice.Connect(); err != nil {
		t.Fatal(err)
	}

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestNXOSDevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	nxosDevice := NXOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	nxosDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestNXOSDevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	nxosDevice := NXOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	_, err := nxosDevice.SendCommand("cmd")
	if err != nil {
		panic(err)
	}

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestNXOSDevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	nxosDevice := NXOSDevice{mockDriver{}, "", &mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	_, err := nxosDevice.SendConfigSet(cmds)
	if err != nil{
		panic(err)
	}

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}
