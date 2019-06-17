package cisco

import "testing"

//type mockDriver struct {
//	ReadReturn   string
//	ReadCount    *int
//	CmdCalls     *string
//	PatternCalls *string
//}
type mockDriver struct {
	ReadSideEffect func() string
	CmdCalls       *string
	PatternCalls   *string
	PromptRegex    *string
}

func (c mockDriver) Connect() {

}
func (c mockDriver) Disconnect() {

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
			println("ddddd1")
			return "switch1"
		case 2:
			println("dddd2")
			return "lorem ipsum 123\nPassword:lorem"
		case 3:
			println("dddd3")
			return "lorem ipsum 123\nswitch#lorem"
		default:
			return ""

		}

	}
	base := CSCODevice{"host", "password", "cisco_asa", "", mockD}
	base.Connect()

	if base.Prompt != "switch1" {
		t.Error("Driver.FindDevicePrompt was not called")
	}
	expected := "#|>, Password:|switch1, switch1, switch1, "

	if patternCalls != expected {
		t.Errorf("wrong Cisco Pattern calls, Expected: (%s) Got: (%s)", expected, patternCalls)
	}

	expected = "enable, password, terminal pager 0, "

	if cmdCalls != expected {
		t.Errorf("wrong Cisco commands calls, Expected: (%s) Got: (%s)", expected, patternCalls)
	}

	expected = "\r?(.*)[#>]"

	if promptRegexCall != expected {
		t.Errorf("wrong Cisco prompt regex calls, Expected: (%s) Got: (%s)", expected, promptRegexCall)
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
			println("ddddd1")
			return "switch1"
		case 2:
			println("dddd2")
			return "lorem ipsum 123\nswitch1#lorem"
		case 3:
			println("dddd3")
			return "switch1#"
		default:
			return ""

		}

	}
	base := CSCODevice{"host", "password", "cisco_asa", "", mockD}
	base.Connect()

	expected := "#|>, Password:|switch1, switch1, "

	if patternCalls2 != expected {
		t.Errorf("wrong Cisco Pattern calls, Expected: (%s) Got: (%s)", expected, patternCalls2)
	}

	expected = "enable, terminal pager 0, "

	if cmdCalls2 != expected {
		t.Errorf("wrong Cisco commands calls, Expected: (%s) Got: (%s)", expected, cmdCalls2)
	}

}

//
//func TestCSCODevice_Disconnect(t *testing.T) {
//
//	base := CSCODevice{"host", "password", "cisco_asa", "", mockDriver{}}
//	base.Disconnect()
//
//	if !stringInSlice("disconnect", calls) {
//		t.Error("Driver.Connect was not called")
//	}
//
//}
//
//func TestCSCODevice_SendCommand(t *testing.T) {
//
//	base := CSCODevice{"host", "password", "cisco_asa", "switch100", mockDriver{}}
//	result, _ := base.SendCommand("cmd")
//	if !strings.Contains(result, "switch100") && !strings.Contains(result, "cmd") {
//		t.Error("Driver.SendCommand was not called as expected")
//	}
//
//}
//
//func TestCSCODevice_SendConfigSet(t *testing.T) {
//	cmds := []string{"cmd1", "cmd2"}
//	base := CSCODevice{"host", "password", "cisco_asa", "switch100", mockDriver{}}
//	result, _ := base.SendConfigSet(cmds)
//
//	if !strings.Contains(result, cmds[0]) && !strings.Contains(result, cmds[1]) && !strings.Contains(result, "switch100") {
//		t.Error("Driver.SendCommandSet was called with wrong args")
//	}
//
//}
//
//func TestASADevice_sessionPreparation(t *testing.T) {
//	base := CSCODevice{"host", "secrect123", "cisco_asa", "switch44", mockDriver{}}
//
//	base.sessionPreparation()
//	println(cmdCalls[0])
//	println(cmdCalls[1])
//	println(cmdCalls[2])
//
//}
//
//func stringInSlice(a string, list []string) bool {
//	for _, b := range list {
//		if b == a {
//			return true
//		}
//	}
//	return false
//}
