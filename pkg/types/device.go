package types

type Device interface {
	OpenSession() error
	Disconnect()
	SendCommand(cmd string) (string, error)
	SendConfigSet(cmds []string) (string, error)
}

type CiscoDevice interface {
	OpenSession() error
	Disconnect()
	SendCommand(cmd string) (string, error)
	SendConfigSet(cmds []string) (string, error)
	SetSecret(secret string)
}