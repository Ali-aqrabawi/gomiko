package driver

import (
	"strings"
	"testing"
	"time"
)

type MockConn struct {
}

func (MockConn) Connect() {

}
func (MockConn) Disconnect() {

}
func (MockConn) Read() (string, error) {
	data := "Loged in as Admin!\n" +
		"loading config\n" +
		"switch1>"
	return data, nil
}

func (MockConn) Write(cmd string) int {
	return 1
}

type MockConn2 struct {
}

func (MockConn2) Connect() {

}
func (MockConn2) Disconnect() {

}

func (MockConn2) Read() (string, error) {
	time.Sleep(6 * time.Second)
	return "", nil
}

func (MockConn2) Write(cmd string) int {
	return 1
}

func TestDriver_FindDevicePrompt(t *testing.T) {

	// [] valid input
	testDriver := Driver{"host", "username", "password", "\n", MockConn{}}
	prompt := testDriver.FindDevicePrompt("\r?(.*)[#>]", ">|#")
	if prompt != "switch1" {
		t.Error("find prompt failed: " + prompt)
	}

	// [] test timeout
	testDriver = Driver{"host", "username", "password", "\n", MockConn2{}}
	go func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
			return
		}()

		testDriver.FindDevicePrompt("\r?(.*)[#>]", ">|#")

	}()

}

func TestDriver_ReadUntil(t *testing.T) {
	testDriver := Driver{"host", "username", "password", "\n", MockConn{}}

	out := testDriver.ReadUntil("switch1")
	if !strings.Contains(out, "Loged in as Admin!") && !strings.Contains(out, "switch1") {
		t.Error("ReadUntil did not return expected data")
	}
}

func TestDriver_SendCommand(t *testing.T) {
	testDriver := Driver{"host", "username", "password", "\n", MockConn{}}

	result, _ := testDriver.SendCommand("show run", "switch1")
	if !strings.Contains(result, "Loged in as Admin!") && !strings.Contains(result, "switch1") {
		t.Error("SendCommand did not return expected data")
	}
}

func TestDriver_SendCommandsSet(t *testing.T) {
	testDriver := Driver{"host", "username", "password", "\n", MockConn{}}

	cmds := []string{"cmd1", "cmd2"}

	result, _ := testDriver.SendCommandsSet(cmds, "switch1")
	if strings.Count(result, "Loged in as Admin!") != 2 {
		t.Error("SendCommandSet did not return expected data")
	}
}
