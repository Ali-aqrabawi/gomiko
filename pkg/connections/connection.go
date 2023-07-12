package connections

import (
	"errors"
)

type Connection interface {
	Connect() error
	Disconnect()
	Read() (string, error)
	Write(cmd string) int
	SetTimeout(timeout uint8)
}

func NewConnection(host string, username string, password string, protocol string, port uint8) (Connection, error) {
	switch protocol {
	case "ssh":
		conn, err := NewSSHConn(host, username, password, port)
		if err != nil {
			return nil, err
		}
		return &conn, nil
	default:
		return nil, errors.New("unsupported protocol: " + protocol)
	}
}
