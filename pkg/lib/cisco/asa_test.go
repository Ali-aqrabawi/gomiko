package cisco

import (
	"testing"
)

type mockBase struct {
	Calls *string
}

func (b mockBase) Connect() {
	*b.Calls = "Connect"

}
func (b mockBase) Disconnect() {
	*b.Calls = "Disconnect"

}
func (b mockBase) SendCommand(cmd string) (string, error) {
	*b.Calls = "SendCommand"

	return "", nil

}
func (b mockBase) SendConfigSet(cmds []string) (string, error) {
	*b.Calls = "SendConfigSet"
	return "", nil

}

func (b mockBase) sessionPreparation() {
	*b.Calls = "sessionPreparation"

}

func TestASADevice_Connect(t *testing.T) {

	// [1] test happy scenario with login -> userMode -> enableMode
	mockb := mockBase{}
	asaDevice := ASADevice{"host", "username", "password", &mockb}

	var calls string
	mockb.Calls = &calls
	asaDevice.Connect()

	if calls != "Connect" {
		t.Error("base.Connect() was not called")
	}

}

func TestASADevice_Disconnect(t *testing.T) {
	mockb := mockBase{}
	asaDevice := ASADevice{"host", "username", "password", &mockb}

	var calls string
	mockb.Calls = &calls
	asaDevice.Disconnect()

	if calls != "Disconnect" {
		t.Error("base.Disconnect() was not called")
	}

}

func TestASADevice_SendCommand(t *testing.T) {
	mockb := mockBase{}
	asaDevice := ASADevice{"host", "username", "password", &mockb}

	var calls string
	mockb.Calls = &calls
	asaDevice.SendCommand("cmd")

	if calls != "SendCommand" {
		t.Error("base.SendCommand() was not called")
	}

}

func TestASADevice_SendConfigSet(t *testing.T) {

	mockb := mockBase{}
	asaDevice := ASADevice{"host", "username", "password", &mockb}

	var calls string
	mockb.Calls = &calls
	cmds := []string{"cmd1", "cmd2"}
	asaDevice.SendConfigSet(cmds)

	if calls != "SendConfigSet" {
		t.Error("base.SendConfigSet() was not called")
	}
}

//func TestCSCODevice_Connect_noUserMode(t *testing.T) {
//
//	// [2] test no userMode scenario login -> enableMode
//	mockD := mockDriver{}
//	var cmdCalls2, patternCalls2, promptRegexCall2 string
//	mockD.CmdCalls = &cmdCalls2
//	mockD.PatternCalls = &patternCalls2
//	mockD.PromptRegex = &promptRegexCall2
//
//	callsCount := 0
//	mockD.ReadSideEffect = func() string {
//		callsCount += 1
//		switch callsCount {
//		case 1:
//			return "switch1"
//		case 2:
//			return "lorem ipsum 123\nswitch1#lorem"
//		case 3:
//			return "switch1#"
//		default:
//			return ""
//
//		}
//
//	}
//	base := CSCODevice{"host", "password", "cisco_ios", "", mockD}
//	base.Connect()
//
//	expected := "#|>, Password:|switch1, switch1, "
//
//	if patternCalls2 != expected {
//		t.Errorf("wrong Cisco Pattern calls, Expected: (%s) Got: (%s)", expected, patternCalls2)
//	}
//
//	expected = "enable, terminal len 0, "
//
//	if cmdCalls2 != expected {
//		t.Errorf("wrong Cisco commands calls, Expected: (%s) Got: (%s)", expected, cmdCalls2)
//	}
//
//}
//
////
//func TestCSCODevice_Disconnect(t *testing.T) {
//	mockD := mockDriver{}
//	var genericCalls string
//	mockD.GenericCalls = &genericCalls
//
//	base := CSCODevice{"host", "password", "cisco_ios", "", mockD}
//
//	base.Disconnect()
//
//	if genericCalls != "disconnect" {
//		t.Error("Driver.Connect was not called")
//	}
//
//}
//
//func TestCSCODevice_SendCommand(t *testing.T) {
//	mockD := mockDriver{}
//	var cmdCalls, patternCalls, promptRegexCall string
//	mockD.CmdCalls = &cmdCalls
//	mockD.PatternCalls = &patternCalls
//	mockD.PromptRegex = &promptRegexCall
//	mockD.ReadSideEffect = func() string {
//		return "show vlans\n" +
//			"vlan 1 v2\n" +
//			"vlan 2 v2 \n" +
//			"switch1# "
//
//	}
//
//	base := CSCODevice{"host", "password", "cisco_ios", "switch1", mockD}
//	result, _ := base.SendCommand("show vlan")
//
//	if !strings.Contains(result, "vlan 1 v2") && !strings.Contains(result, "vlan 2 v2 ") {
//		t.Error("wrong result returned")
//	}
//
//	expected := "show vlan, "
//
//	if cmdCalls != expected {
//		t.Errorf("wrong commands calls, expected: (%s) got: (%s)", expected, cmdCalls)
//	}
//
//}
//
//func TestCSCODevice_SendConfigSet(t *testing.T) {
//	mockD := mockDriver{}
//	var cmdCalls, patternCalls, promptRegexCall string
//	mockD.CmdCalls = &cmdCalls
//	mockD.PatternCalls = &patternCalls
//	mockD.PromptRegex = &promptRegexCall
//	mockD.ReadSideEffect = func() string {
//		return "show vlans\n" +
//			"vlan 1 v2\n" +
//			"vlan 2 v2 \n" +
//			"switch1# "
//
//	}
//
//	base := CSCODevice{"host", "password", "cisco_ios", "switch1", mockD}
//	cmds := []string{"show interface", "show ip route"}
//	base.SendConfigSet(cmds)
//
//	expected := "config term, show interface, show ip route, end, "
//
//	if cmdCalls != expected {
//		t.Errorf("wrong commands calls, expected: (%s) got: (%s)", expected, cmdCalls)
//	}
//
//}
