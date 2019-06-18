package driver

import (
	"errors"
	"github.com/Ali-aqrabawi/gomiko/pkg/connections"
	"github.com/Ali-aqrabawi/gomiko/pkg/utils"
	"regexp"
	"strings"
	"time"
)

type Driver struct {
	Host       string
	Username   string
	Password   string
	Return     string `defaults:"\n"`
	connection connections.Connection
}

func (d *Driver) Connect() {
	utils.LogInfo(d.Host, "connecting to Device...")
	d.connection.Connect()

}

func (d *Driver) Disconnect() {
	utils.LogInfo(d.Host, "disconnecting Device...")
	d.connection.Disconnect()

}

func (d *Driver) SendCommand(cmd string, expectPattern string) (string, error) {
	utils.LogInfo(d.Host, "sending command: "+cmd)

	if d.connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	cmd += d.Return

	d.connection.Write(cmd)

	result := d.ReadUntil(expectPattern)

	return result, nil

}

func (d *Driver) SendCommandsSet(cmds []string, expectPattern string) (string, error) {
	utils.LogInfo(d.Host, "sending config set: "+strings.Join(cmds, ", "))
	if d.connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}
	var results string

	for _, cmd := range cmds {
		out, _ := d.SendCommand(cmd, expectPattern)
		results += out
	}
	return results, nil

}

func (d *Driver) FindDevicePrompt(regex string, pattern string) string {
	var out string
	r, _ := regexp.Compile(regex)

	if pattern != "" {
		out = d.ReadUntil(pattern)
	} else {
		out, _ = d.connection.Read()
	}
	if !r.MatchString(out) {
		utils.LogFatal("", "failed to find prompt, pattern: "+pattern+" , output: "+out, nil)
	}
	return r.FindStringSubmatch(out)[1]

}

func (d *Driver) ReadUntil(pattern string) string {
	outputChan := make(chan string)

	go func(d *Driver, pattern string) {
		buffChan := make(chan string)
		go readRoutine(d, pattern, buffChan)
		select {
		case recv := <-buffChan:
			outputChan <- recv

		case <-time.After(time.Duration(4) * time.Second):
			panic("timeout while reading, read pattern not found pattern: " + pattern)

		}

	}(d, pattern)

	return <-outputChan

}


func readRoutine(d *Driver, pattern string, buffChan chan<- string) {
	var result string
	result, err := d.connection.Read()

	r, _ := regexp.Compile(pattern)

	for (err == nil) && (!r.MatchString(result)) {
		outSlice, _ := d.connection.Read()
		result += outSlice

	}
	//for (err == nil) && (!strings.Contains(result, pattern)) {
	//	outSlice, _ := d.Connection.Read()
	//	result += outSlice
	//
	//}

	buffChan <- result

}

