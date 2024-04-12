package tty2web

import (
	"bytes"
	"os/exec"
)

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
