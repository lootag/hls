package hls

import (
	"bytes"
	"os/exec"
)

type command struct {
	app       string
	arguments []string
}

func (this *command) constructor(app string, arguments []string) *command {
	this.app = app
	this.arguments = arguments
	return this
}

func NewCommand(app string, arguments []string) *command {
	command := &command{}
	return command.constructor(app, arguments)
}

func (this *command) App() string {
	return this.app
}

func (this *command) Arguments() []string {
	return this.arguments
}

func (this *command) Execute() {
	cmd := exec.Command(this.app, this.arguments...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		panic(string(stderr.Bytes()))
	}
}
