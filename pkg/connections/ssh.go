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
	client *ssh.Client
	reader io.Reader
	writer io.WriteCloser
}

func NewSSHConn(hostname string, username string, password string, port uint8) (SSHConn, error) {
	sshConn := SSHConn{}
	interactive := getInteractiveCallBack(password)
	sshConfig := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password), ssh.KeyboardInteractive(interactive)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         6 * time.Second,
	}
	sshConfig.Ciphers = append(sshConfig.Ciphers, ciphers...)
	addr := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {

		return sshConn, errors.New("failed to connect to device: " + err.Error())
	}

	sshConn.client = conn

	return sshConn, nil

}

func NewSSHConnFromClient(client *ssh.Client) (SSHConn, error) {
	sshConn := SSHConn{}
	if client.Conn == nil {
		return sshConn, errors.New("*ssh.Client has no Conn, make sure to call .Dial() before")
	}
	sshConn.client = client
	return sshConn, nil

}

func (c *SSHConn) OpenSession() error {

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

	buff := make([]byte, 2048)

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
		for n, _ := range questions {
			answers[n] = password
		}

		return answers, nil
	}

}
