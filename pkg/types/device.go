package types

type Device interface {
	Connect()
	Disconnect()
	SendCommand(cmd string) (string, error)
	SendConfigSet(cmds []string) (string, error)
}


