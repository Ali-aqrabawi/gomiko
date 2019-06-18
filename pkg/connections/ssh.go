package connections

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/utils"
	"golang.org/x/crypto/ssh"
	"io"
)

var ciphers = []string{"3des-cbc", "aes128-cbc", "aes192-cbc", "aes256-cbc", "aes128-ctr"}

type SSHConn struct {
	Host     string
	Username string
	Password string
	client   *ssh.Client
	reader   io.Reader
	writer   io.WriteCloser
}

func NewSSHConn(host string, username string, password string) *SSHConn {
	return &SSHConn{host, username, password, nil, nil, nil}
}

func (c *SSHConn) Connect() {

	sshConfig := &ssh.ClientConfig{User: c.Username, Auth: []ssh.AuthMethod{ssh.Password(c.Password)}, HostKeyCallback: ssh.InsecureIgnoreHostKey(), Timeout: 0}
	sshConfig.Ciphers = append(sshConfig.Ciphers, ciphers...)
	addr := c.Host + ":22"
	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		utils.LogFatal(c.Host, "failed to connect: ", err)
	}

	session, err := conn.NewSession()

	if err != nil {
		utils.LogFatal(c.Host, "failed to open session: ", err)
	}

	reader, _ := session.StdoutPipe()
	writer, _ := session.StdinPipe()

	c.client = conn
	c.reader = reader
	c.writer = writer

	if err != nil {
		panic(err)
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("vt100", 0, 200, modes); err != nil {

		utils.LogFatal(c.Host, "failed to request pty: ", err)
	}
	if err := session.Shell(); err != nil {
		utils.LogFatal(c.Host, "failed to invoke shell: ", err)
	}

}

func (c *SSHConn) Disconnect() {

	err := c.client.Close()
	if err != nil {
		utils.LogFatal(c.Host, "failed to disconnect: ", err)
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
