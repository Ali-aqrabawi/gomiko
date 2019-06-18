package mikrotik

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

func (c mockDriver) Connect() {

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

	return c.ReadUntil(expectPattern), nil

}
func (c mockDriver) SendCommandsSet(cmds []string, expectPattern string) (string, error) {
	for _, cmd := range cmds {
		c.SendCommand(cmd, expectPattern)
	}
	return c.ReadUntil(expectPattern), nil

}

func (c mockDriver) FindDevicePrompt(regex string, pattern string) string {
	*c.PromptRegex = regex
	return c.ReadUntil(pattern)

}

func (c mockDriver) ReadUntil(pattern string) string {
	*c.PatternCalls += pattern + ", "

	return c.ReadSideEffect()

}


func TestMikroTikROS_Connect(t *testing.T) {

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
			return "@MikroTik] >"
		case 2:
			return "lorem ipsum 123\n@MikroTik] >lorem"
		case 3:
			return "lorem ipsum 123\n@MikroTik] >lorem"
		default:
			return ""

		}

	}
	base := MikroTikROS{"host", "password", "mikrotik", "", mockD}
	base.Connect()

	if base.Prompt != "@MikroTik] >" {
		t.Error("Driver.FindDevicePrompt was not called")
	}
	expected := "] >, @MikroTik] >, "

	if patternCalls != expected {
		t.Errorf("wrong Mikrotik Pattern calls, Expected: (%s) Got: (%s)", expected, patternCalls)
	}

	expected = ", "  // MikroTik does not need any sessionPreparation

	if cmdCalls != expected {
		t.Errorf("wrong Mikrotik commands calls, Expected: (%s) Got: (%s)", expected, cmdCalls)
	}

	expected = "\\[.*(@.*\\] >)"

	if promptRegexCall != expected {
		t.Errorf("wrong Mikrotik prompt regex calls, Expected: (%s) Got: (%s)", expected, promptRegexCall)
	}

}


func TestMikroTikROS_Disconnect(t *testing.T) {
	mockD := mockDriver{}
	var genericCalls string
	mockD.GenericCalls = &genericCalls

	base := MikroTikROS{"host", "password", "mikrotik", "", mockD}

	base.Disconnect()

	if genericCalls != "disconnect" {
		t.Error("Driver.Connect was not called")
	}

}

func TestMikroTikROS_SendCommand(t *testing.T) {
	mockD := mockDriver{}
	var cmdCalls, patternCalls, promptRegexCall string
	mockD.CmdCalls = &cmdCalls
	mockD.PatternCalls = &patternCalls
	mockD.PromptRegex = &promptRegexCall
	mockD.ReadSideEffect = func() string {
		return "ip route print\n" +
			   "0 ADS  0.0.0.0/0\n" +
			   "1 ADC  192.168.122.0/24\n" +
			   "@MikroTik] > "

	}

	base := MikroTikROS{"host", "password", "mikrotik", "@MikroTik] >", mockD}
	result, _ := base.SendCommand("ip route print")

	if !strings.Contains(result, "0 ADS  0.0.0.0/0") &&
		!strings.Contains(result, "1 ADC  192.168.122.0/24") {
		t.Error("wrong result returned")
	}

	expected := "ip route print, "

	if cmdCalls != expected {
		t.Errorf("wrong commands calls, expected: (%s) got: (%s)", expected, cmdCalls)
	}

}

func TestMikroTikROS_SendConfigSet(t *testing.T) {
	mockD := mockDriver{}
	var cmdCalls, patternCalls, promptRegexCall string
	mockD.CmdCalls = &cmdCalls
	mockD.PatternCalls = &patternCalls
	mockD.PromptRegex = &promptRegexCall
	mockD.ReadSideEffect = func() string {
		return "ip route print\n" +
			   "0 ADS  0.0.0.0/0\n" +
			   "1 ADC  192.168.122.0/24\n" +
			   "@MikroTik] > "

	}

	base := MikroTikROS{"host", "password", "cisco_ios", "switch1", mockD}
	cmds := []string{"mikrotik command1", "mikrotik command2"}
	base.SendConfigSet(cmds)

	expected := "mikrotik command1, mikrotik command2, "

	if cmdCalls != expected {
		t.Errorf("wrong commands calls, expected: (%s) got: (%s)", expected, cmdCalls)
	}

}
