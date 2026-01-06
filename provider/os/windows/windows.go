package windows

import "os/exec"

type Windows struct {
}

func (Windows) Provider() string { return "windows" }

func (Windows) GetHostname() (string, error) {
	cmd := exec.Command("hostname")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
