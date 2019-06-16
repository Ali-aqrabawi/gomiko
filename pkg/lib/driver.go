package lib

import (
	"errors"
	"gomiko/pkg/connections"
	"regexp"
	"strings"
	"time"
)

type Driver struct {
	Connection connections.Connection
	Return     string `defaults:"\n"`
}

func (d Driver) SendCommand(cmd string, expectPattern string) (string, error) {
	if d.Connection == nil {
		return "", errors.New("not connected to device, make sure to call .Connect() first")
	}

	cmd += d.Return

	d.Connection.Write(cmd)

	result := d.ReadUntil(expectPattern)

	return result, nil

}

func (d Driver) SendCommandsSet(cmds []string, expectPattern string) (string, error) {
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

func (d Driver) FindDevicePrompt(regex string, pattern string) string {
	var out string
	r, _ := regexp.Compile(regex)
	d.Connection.Write(d.Return)
	if pattern != "" {
		out = d.ReadUntil(pattern)
	} else {
		out, _ = d.Connection.Read()
	}

	return r.FindStringSubmatch(out)[1]

}

func readRoutine(d Driver, pattern string, buffChan chan<- string) {
	var result string
	result, err := d.Connection.Read()

	for (err == nil) && (!strings.Contains(result, pattern)) {
		outSlice, _ := d.Connection.Read()
		result += outSlice

	}

	buffChan <- result

}

func (d Driver) ReadUntil(pattern string) string {
	outputChan := make(chan string)

	go func(d Driver, pattern string) {
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
