package connections

type Connection interface {
	Connect()
	Disconnect()
	Read() (string, error)
	Write(cmd string) int
}

func NewConnection(host string, username string, password string, protocol string) Connection {
	switch protocol {
	case "ssh":
		return NewSSHConn(host, username, password)

	default:
		panic("invalid protocol: " + protocol)

	}

}
