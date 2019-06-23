package connections

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
)

type SSHClientConn struct {
	client *ssh.Client
	reader io.Reader
	writer io.WriteCloser
}

func Create(client *ssh.Client) *SSHClientConn {
	return &SSHClientConn{client, nil, nil}
}

func (c *SSHClientConn) Connect() error {

	session, err := c.client.NewSession()

	if err != nil {
		return errors.New("failed to Start a new session: " + err.Error())
	}

	reader, _ := session.StdoutPipe()
	writer, _ := session.StdinPipe()

	c.reader = reader
	c.writer = writer

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("vt100", 0, 200, modes); err != nil {

		return errors.New("failed to request Pty: " + err.Error())
	}
	if err := session.Shell(); err != nil {
		return errors.New("failed to invoke shell: " + err.Error())
	}
	return nil

}

func (c *SSHClientConn) Disconnect() {

	err := c.client.Close()
	if err != nil {
		log.Println("warning, device close failed: ", err)
	}

}

func (c *SSHClientConn) Read() (string, error) {

	buff := make([]byte, 2048)

	n, err := c.reader.Read(buff)
	return string(buff[:n]), err

}

func (c *SSHClientConn) Write(cmd string) int {

	commandBytes := []byte(cmd)
	code, _ := c.writer.Write(commandBytes)
	return code

}
