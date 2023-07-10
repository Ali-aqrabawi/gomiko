package connections

import (
	"errors"
)

type Connection interface {
	Connect() error
	Disconnect()
	Read() (string, error)
	Write(cmd string) int
}

func NewConnection(host string, username string, password string, protocol string, port uint8, timeout ...uint8) (Connection, error) {
	switch protocol {
	case "ssh":
		if len(timeout) > 0 {
			conn, err := NewSSHConn(host, username, password, port, timeout[0])
			if err != nil {
				return nil, err
			}
			return &conn, nil
		} else {
			conn, err := NewSSHConn(host, username, password, port)
			if err != nil {
				return nil, err
			}
			return &conn, nil
		}
	default:
		return nil, errors.New("unsupported protocol: " + protocol)
	}
}
