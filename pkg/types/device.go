package types

type Device interface {
	Connect() error
	Disconnect()
	SendCommand(cmd string) (string, error)
	SendConfigSet(cmds []string) (string, error)
}


