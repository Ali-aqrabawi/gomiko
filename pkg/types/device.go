package types

type Device interface {
	Connect() error
	Disconnect()
	SendCommand(cmd string) (string, error)
	SendConfigSet(cmds []string) (string, error)
	SetTimeout(timeout uint8)
}

type CiscoDevice interface {
	Connect() error
	Disconnect()
	SendCommand(cmd string) (string, error)
	SendConfigSet(cmds []string) (string, error)
	SetSecret(secret string)
	SetTimeout(timeout uint8)
}
