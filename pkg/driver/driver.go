package driver

import (
	"errors"
	"github.com/Ali-aqrabawi/gomiko/pkg/connections"
	"regexp"
	"time"
)

type Driver struct {
	Connection connections.Connection
	Return     string `defaults:"\n"`
}

func (d *Driver) Connect() error {
	err := d.Connection.Connect()
	return err

}

func (d *Driver) Disconnect() {
	d.Connection.Disconnect()

}

func (d *Driver) SendCommand(cmd string, expectPattern string) (string, error) {
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	cmd += d.Return

	d.Connection.Write(cmd)

	result, err := d.ReadUntil(expectPattern)

	return result, err

}

func (d *Driver) SendCommandsSet(cmds []string, expectPattern string) (string, error) {
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}
	var results string

	for _, cmd := range cmds {
		out, _ := d.SendCommand(cmd, expectPattern)
		results += out
	}
	return results, nil

}

func (d *Driver) FindDevicePrompt(regex string, pattern string) (string, error) {
	var out string
	var err error
	r, _ := regexp.Compile(regex)

	if pattern != "" {
		out, err = d.ReadUntil(pattern)
		if err != nil {
			return "", err
		}
	} else {
		out, _ = d.Connection.Read()
	}
	if !r.MatchString(out) {
		return "", errors.New("failed to find prompt, pattern: " + pattern + " , output: " + out)
	}
	return r.FindStringSubmatch(out)[1], nil

}

func (d *Driver) ReadUntil(pattern string) (string, error) {

	outputChan := make(chan string)
	var err error

	go func(d *Driver, pattern string) {
		buffChan := make(chan string)
		go readRoutine(d, pattern, buffChan)
		select {
		case recv := <-buffChan:
			outputChan <- recv

		case <-time.After(time.Duration(4) * time.Second):
			err = errors.New("timeout while reading, read pattern not found pattern: " + pattern)
			close(outputChan)
		}

	}(d, pattern)


	return <-outputChan, err

}

func (d Driver) SetTimeout(timeout int) {
	d.Connection.SetTimeout(timeout)
}

func (d Driver) SetReturn(Return string) {
	d.Return = Return
}

func readRoutine(d *Driver, pattern string, buffChan chan<- string) {
	var result string
	result, err := d.Connection.Read()

	r, _ := regexp.Compile(pattern)

	for (err == nil) && (!r.MatchString(result)) {
		outSlice, _ := d.Connection.Read()
		result += outSlice

	}
	//for (err == nil) && (!strings.Contains(result, pattern)) {
	//	outSlice, _ := d.Connection.Read()
	//	result += outSlice
	//
	//}

	buffChan <- result

}
