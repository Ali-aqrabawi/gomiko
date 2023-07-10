package connections

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"time"
)

var ciphers = []string{
	"aes256-ctr",
	"aes128-ctr",
	"aes128-cbc",
	"3des-cbc",
	"aes192-ctr",
	"aes192-cbc",
	"aes256-cbc",
	"aes128-gcm@openssh.com"}

type SSHConn struct {
	addr     string
	username string
	password string
	client   *ssh.Client
	reader   io.Reader
	writer   io.WriteCloser
	timeout  uint8
}

func NewSSHConn(hostname string, username string, password string, port uint8, timeout ...uint8) (SSHConn, error) {
	sshConn := SSHConn{}

	addr := fmt.Sprintf("%s:%d", hostname, port)
	sshConn.addr = addr
	sshConn.username = username
	sshConn.password = password
	if len(timeout) > 0 {
		sshConn.timeout = timeout[0]
	} else {
		sshConn.timeout = 6 // Default timeout is 6 seconds
	}

	return sshConn, nil

}

func (c *SSHConn) Connect() error {
	interactive := getInteractiveCallBack(c.password)
	sshConfig := &ssh.ClientConfig{
		User:            c.username,
		Auth:            []ssh.AuthMethod{ssh.Password(c.password), ssh.KeyboardInteractive(interactive)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(c.timeout) * time.Second,
	}
	sshConfig.Ciphers = append(sshConfig.Ciphers, ciphers...)
	conn, err := ssh.Dial("tcp", c.addr, sshConfig)
	if err != nil {

		return errors.New("failed to connect to device: " + err.Error())
	}
	c.client = conn

	session, err := c.client.NewSession()

	if err != nil {
		return errors.New("failed to Start a new session: " + err.Error())
	}

	reader, _ := session.StdoutPipe()
	writer, _ := session.StdinPipe()

	//c.client = conn
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

func (c *SSHConn) Disconnect() {

	err := c.client.Close()
	if err != nil {
		log.Println("warning, device close failed: ", err)
	}

}

func (c *SSHConn) Read() (string, error) {

	buff := make([]byte, 204800)

	n, err := c.reader.Read(buff)
	return string(buff[:n]), err

}

func (c *SSHConn) Write(cmd string) int {

	commandBytes := []byte(cmd)
	code, _ := c.writer.Write(commandBytes)
	return code

}

func getInteractiveCallBack(password string) ssh.KeyboardInteractiveChallenge {

	return func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
		answers = make([]string, len(questions))
		// The second parameter is unused
		for n := range questions {
			answers[n] = password
		}

		return answers, nil
	}

}
