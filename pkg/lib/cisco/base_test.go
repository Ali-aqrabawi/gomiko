package cisco

import (
	"strings"
	"testing"
)

type mockDriver struct {
	ReadSideEffect func() string
	CmdCalls       *string
	PatternCalls   *string
	PromptRegex    *string
	GenericCalls   *string
}

func (c mockDriver) Connect() error {
	return nil

}
func (c mockDriver) Disconnect() {
	*c.GenericCalls = "disconnect"

}
func (c mockDriver) SendCommand(cmd string, expectPattern string) (string, error) {
	if c.CmdCalls != nil {
		*c.CmdCalls += cmd + ", "
	} else {
		c.CmdCalls = &cmd
	}

	return c.ReadUntil(expectPattern)

}
func (c mockDriver) SendCommandsSet(cmds []string, expectPattern string) (string, error) {
	for _, cmd := range cmds {
		_, err := c.SendCommand(cmd, expectPattern)
		if err != nil {
			panic(err)
		}
	}
	return c.ReadUntil(expectPattern)

}

func (c mockDriver) FindDevicePrompt(regex string, pattern string) (string, error) {
	*c.PromptRegex = regex
	return c.ReadUntil(pattern)

}

func (c mockDriver) ReadUntil(pattern string) (string, error) {
	*c.PatternCalls += pattern + ", "

	return c.ReadSideEffect(), nil

}

func (c mockDriver) SetTimeout(timeout uint8) {
}

func TestCSCODevice_Connect_userMode(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode
	mockD := mockDriver{}
	var cmdCalls, patternCalls, promptRegexCall string
	mockD.CmdCalls = &cmdCalls
	mockD.PatternCalls = &patternCalls
	mockD.PromptRegex = &promptRegexCall

	callsCount := 0
	mockD.ReadSideEffect = func() string {
		callsCount += 1
		switch callsCount {
		case 1:
			return "switch1"
		case 2:
			return "lorem ipsum 123\nPassword:lorem"
		case 3:
			return "lorem ipsum 123\nswitch#lorem"
		default:
			return ""
		}
	}

	base := CSCODevice{mockD, "", "", ""}
	if err := base.Connect(); err != nil {
		t.Fatal(err)
	}

	if base.Prompt != "switch1" {
		t.Error("Driver.FindDevicePrompt was not called")
	}
	expected := "#|>, Password:|switch1, switch1, switch1, "

	if patternCalls != expected {
		t.Errorf("wrong Cisco Pattern calls, Expected: (%s) Got: (%s)", expected, patternCalls)
	}

	expected = "enable, , terminal len 0, "

	// Test no Secret
	if cmdCalls != expected {
		t.Errorf("wrong Cisco commands calls, Expected: (%s) Got: (%s)", expected, cmdCalls)
	}

	// expected = "\r?(.*)[#>]" - This does not match what's used in the Connect function making the call to FindDevicePrompt. What's below is used instead.
	expected = "\r\n?(\\S+)[#>]"

	if promptRegexCall != expected {
		t.Errorf("wrong Cisco prompt regex calls, Expected: (%s) Got: (%s)", expected, promptRegexCall)
	}

	// Test with Secret
	base.SetSecret("mySecret")
	cmdCalls = ""
	callsCount = 0
	if err := base.Connect(); err != nil {
		t.Fatal(err)
	}
	expected = "enable, mySecret, terminal len 0, "
	if cmdCalls != expected {
		t.Errorf("wrong Cisco commands calls, Expected: (%s) Got: (%s)", expected, cmdCalls)
	}

}

func TestCSCODevice_Connect_noUserMode(t *testing.T) {

	// [2] test no userMode scenario login -> enableMode
	mockD := mockDriver{}
	var cmdCalls2, patternCalls2, promptRegexCall2 string
	mockD.CmdCalls = &cmdCalls2
	mockD.PatternCalls = &patternCalls2
	mockD.PromptRegex = &promptRegexCall2

	callsCount := 0
	mockD.ReadSideEffect = func() string {
		callsCount += 1
		switch callsCount {
		case 1:
			return "switch1"
		case 2:
			return "lorem ipsum 123\nswitch1#lorem"
		case 3:
			return "switch1#"
		default:
			return ""

		}

	}
	base := CSCODevice{mockD, "cisco_ios", "", ""}
	if err := base.Connect(); err != nil {
		t.Fatal(err)
	}

	expected := "#|>, Password:|switch1, switch1, "

	if patternCalls2 != expected {
		t.Errorf("wrong Cisco Pattern calls, Expected: (%s) Got: (%s)", expected, patternCalls2)
	}

	expected = "enable, terminal len 0, "

	if cmdCalls2 != expected {
		t.Errorf("wrong Cisco commands calls, Expected: (%s) Got: (%s)", expected, cmdCalls2)
	}

}

func TestCSCODevice_Disconnect(t *testing.T) {
	mockD := mockDriver{}
	var genericCalls string
	mockD.GenericCalls = &genericCalls

	base := CSCODevice{mockD, "cisco_ios", "", ""}

	base.Disconnect()

	if genericCalls != "disconnect" {
		t.Error("Driver.Disconnect() was not called")
	}

}

func TestCSCODevice_SendCommand(t *testing.T) {
	mockD := mockDriver{}
	var cmdCalls, patternCalls, promptRegexCall string
	mockD.CmdCalls = &cmdCalls
	mockD.PatternCalls = &patternCalls
	mockD.PromptRegex = &promptRegexCall
	mockD.ReadSideEffect = func() string {
		return "show vlans\n" +
			"vlan 1 v2\n" +
			"vlan 2 v2 \n" +
			"switch1# "

	}

	base := CSCODevice{mockD, "cisco_ios", "switch1", ""}
	result, _ := base.SendCommand("show vlan")

	if !strings.Contains(result, "vlan 1 v2") && !strings.Contains(result, "vlan 2 v2 ") {
		t.Error("wrong result returned")
	}

	expected := "show vlan, "

	if cmdCalls != expected {
		t.Errorf("wrong commands calls, expected: (%s) got: (%s)", expected, cmdCalls)
	}

}

func TestCSCODevice_SendConfigSet(t *testing.T) {
	mockD := mockDriver{}
	var cmdCalls, patternCalls, promptRegexCall string
	mockD.CmdCalls = &cmdCalls
	mockD.PatternCalls = &patternCalls
	mockD.PromptRegex = &promptRegexCall
	mockD.ReadSideEffect = func() string {
		return "show vlans\n" +
			"vlan 1 v2\n" +
			"vlan 2 v2 \n" +
			"switch1# "

	}

	base := CSCODevice{mockD, "cisco_ios", "switch1", ""}
	cmds := []string{"show interface", "show ip route"}
	_, err := base.SendConfigSet(cmds)
	if err != nil {
		panic(err)
	}

	expected := "config term, show interface, show ip route, end, "

	if cmdCalls != expected {
		t.Errorf("wrong commands calls, expected: (%s) got: (%s)", expected, cmdCalls)
	}

}
