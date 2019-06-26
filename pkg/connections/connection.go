package connections

import (
	"errors"
	"golang.org/x/crypto/ssh"
)

type Connection interface {
	OpenSession() error
	Disconnect()
	Read() (string, error)
	Write(cmd string) int
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

func NewConnectionFromClient(client *ssh.Client) (Connection, error) {
	conn, err := NewSSHConnFromClient(client)
	if err != nil {
		return nil, err
	}
	return &conn, nil

}
