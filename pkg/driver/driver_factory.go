package driver

import "github.com/Ali-aqrabawi/gomiko/pkg/connections"

type IDriver interface {
	Connect() error
	Disconnect()
	SendCommand(cmd string, expectPattern string) (string, error)
	SendCommandsSet(cmds []string, expectPattern string) (string, error)
	FindDevicePrompt(regex string, pattern string) (string, error)
	ReadUntil(pattern string) (string, error)
}

func NewDriver(Host string, Username string, Password string, Return string, protocol string) IDriver {
	connection := connections.NewConnection(Host, Username, Password, protocol)
	return &Driver{Host, Username, Password, Return, connection}

}
