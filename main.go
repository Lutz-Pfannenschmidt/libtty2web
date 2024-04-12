package libtty2web

import (
	"os/exec"
	"strings"
)

type Tty2Web struct {
	options []option
	binary  string
	process *exec.Cmd
	command string
}

func NewTty2Web(command string, options ...option) *Tty2Web {
	return &Tty2Web{options: options, binary: "tty2web", command: command}
}

func (t *Tty2Web) AddOptions(option ...option) {
	t.options = append(t.options, option...)
}

func (t *Tty2Web) SetBinary(binary string) {
	t.binary = binary
}

func (t *Tty2Web) GetCommand() string {
	return t.binary + " " + strings.Join(t.buildCommand(), " ")
}

func (t *Tty2Web) buildCommand() []string {
	command := []string{}
	for _, arg := range t.options {
		if arg.hasValue {
			command = append(command, "--"+arg.name, arg.value)
		} else {
			command = append(command, "--"+arg.name)
		}
	}
	command = append(command, t.command)
	return command
}

func (t *Tty2Web) Run() error {
	t.process = exec.Command(t.command, t.buildCommand()...)
	err := t.process.Run()
	if err != nil {
		return err
	}
	return nil
}

func (t *Tty2Web) Kill() error {
	err := t.process.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}

func (t *Tty2Web) Wait() error {
	err := t.process.Wait()
	if err != nil {
		return err
	}
	return nil
}
