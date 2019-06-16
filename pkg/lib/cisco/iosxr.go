package cisco

type IOSXRDevice struct {
	Host     string
	Username string
	Password string
	base   CSCODevice
}

func (d *IOSXRDevice) Connect() {
	d.base.Connect()

}

func (d *IOSXRDevice) Disconnect() {

	d.base.Disconnect()

}


func (d *IOSXRDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *IOSXRDevice) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)

}
