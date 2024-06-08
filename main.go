package libtty2web

import (
	"os/exec"
	"strings"
)

type Tty2Web struct {
	options []option  // options to be passed to the binary
	binary  string    // binary to be executed
	process *exec.Cmd // process to be executed
	command []string  // command to be executed
}

// NewTty2Web creates a new Tty2Web instance with the given command
func NewTty2Web(command ...string) *Tty2Web {
	return &Tty2Web{options: []option{}, binary: "tty2web", command: command}
}

// AddOptions adds options to the Tty2Web instance
func (t *Tty2Web) AddOptions(option ...option) {
	t.options = append(t.options, option...)
}

// SetBinary sets the binary to be used by the Tty2Web instance
func (t *Tty2Web) SetBinary(binary string) {
	t.binary = binary
}

// GetCommand returns the command to be executed
func (t *Tty2Web) GetCommand() string {
	return t.binary + " " + strings.Join(t.buildCommand(), " ")
}

// buildCommand builds the command to be executed
func (t *Tty2Web) buildCommand() []string {
	command := []string{}
	for _, arg := range t.options {
		if arg.hasValue {
			command = append(command, "--"+arg.name, arg.value)
		} else {
			command = append(command, "--"+arg.name)
		}
	}
	command = append(command, t.command...)
	return command
}

// Run executes the command
func (t *Tty2Web) Run() error {
	t.process = exec.Command(t.binary, t.buildCommand()...)
	err := t.process.Run()
	if err != nil {
		return err
	}
	return nil
}

// Kill kills the process
func (t *Tty2Web) Kill() error {
	err := t.process.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}

// Wait waits for the process to finish
func (t *Tty2Web) Wait() error {
	err := t.process.Wait()
	if err != nil {
		return err
	}
	return nil
}
