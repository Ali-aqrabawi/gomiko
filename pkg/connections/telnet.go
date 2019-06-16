package connections

import (
	"io"
)

type TelnetConn struct {
	Host     string
	Username string
	Password string
	reader   io.Reader
	writer   io.WriteCloser
}

func NewTelnetConn(host string, username string, password string) *SSHConn {
	return &SSHConn{host, username, password, nil, nil, nil}
}

func (c *TelnetConn) Connect() {

}

func (c *TelnetConn) Disconnect() {

}

func (c *TelnetConn) Read() (string, error) {

	return "", nil

}

func (c *TelnetConn) Write(cmd string) int {
	return 0

}
