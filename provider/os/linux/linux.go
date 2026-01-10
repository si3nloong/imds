package linux

import (
	"errors"
	"net"
	"os"
	"os/exec"
	"strings"
)

type Linux struct{}

func (Linux) Provider() string { return "Linux" }

func (Linux) InstanceID() (string, error) {
	// Try to read machine-id
	id, err := os.ReadFile("/etc/machine-id")
	if err == nil {
		return strings.TrimSpace(string(id)), nil
	}
	id, err = os.ReadFile("/var/lib/dbus/machine-id")
	if err == nil {
		return strings.TrimSpace(string(id)), nil
	}
	return "", nil
}

func (Linux) InstanceType() (string, error) {
	return "", nil
}

func (Linux) Region() (string, error) {
	return "", nil
}

func (Linux) Zone() (string, error) {
	return "", nil
}

func (Linux) PublicIP() (string, error) {
	cmd := exec.Command("curl", "ifconfig.me")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (Linux) PrivateIP() (string, error) {
	// Get all network interfaces
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// Check if the address is a network IP address
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// Check if it is an IPv4 address and not link-local (169.254.x.x)
			if ipnet.IP.To4() != nil && !ipnet.IP.IsLinkLocalUnicast() {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("unable to find machine local ip")
}
