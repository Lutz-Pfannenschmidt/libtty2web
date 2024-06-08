package libtty2web

import (
	"bytes"
	"os/exec"
)

// GetHelpMsg returns the help message of the binary
func GetHelpMsg(binary string) (string, error) {

	cmd := exec.Command(binary, "--help")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return "", err
	}

	return out.String(), nil
}

// GetVersion returns the version of the binary
func GetVersion(binary string) (string, error) {

	cmd := exec.Command(binary, "--version")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return "", err
	}

	return out.String(), nil
}
