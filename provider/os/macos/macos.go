package macos

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"net"
	"os/exec"
	"time"
	"unsafe"
)

var metadata Metadata

func init() {
	metadata, _ = GetMetadata()
}

const (
	IOPlatformUUID         = "IOPlatformUUID"
	IOPlatformSerialNumber = "IOPlatformSerialNumber"
	SerialNumber           = "serial-number"
	PlatformName           = "platform-name"
	Model                  = "model"
	Compatible             = "compatible"
	DeviceType             = "device_type"
	RegionInfo             = "region-info"
	Manufacturer           = "manufacturer"
	RegulatoryModelNumber  = "egulatory-model-number"
	Timestamp              = "time-stamp"
)

type Metadata struct {
	PlatformUUID          string
	PlatformName          string
	Model                 string
	DeviceType            string
	ModelNumber           string
	Timestamp             time.Time
	PlatformSerialNumber  string
	SerialNumber          string
	Compatible            string
	Manufacturer          string
	RegionInfo            string
	RegulatoryModelNumber string
}

func GetMetadata() (Metadata, error) {
	// Query ioreg for meta data
	cmd := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice")
	output, err := cmd.Output()
	if err != nil {
		return Metadata{}, err
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(output))
	md := Metadata{}
	for scanner.Scan() {
		lineBytes := scanner.Bytes()
		parts := bytes.Split(lineBytes, []byte("="))
		if len(parts) == 2 {
			key := strip(unsafe.String(unsafe.SliceData(parts[0]), len(parts[0])))
			value := stripValueBytes(parts[1])
			switch key {
			case IOPlatformUUID:
				md.PlatformUUID = value
			case IOPlatformSerialNumber:
				md.PlatformSerialNumber = value
			case Model:
				md.Model = value
			case Timestamp:
				// md.Timestamp = value
			case DeviceType:
				md.DeviceType = value
			case Compatible:
				md.Compatible = value
			case RegulatoryModelNumber:
				md.RegulatoryModelNumber = value
			case RegionInfo:
				b, err := hex.DecodeString(value)
				if err == nil {
					md.RegionInfo = string(bytes.TrimSpace(b))
				}
			case SerialNumber:
				b, err := hex.DecodeString(value)
				if err == nil {
					md.SerialNumber = string(b)
				}
			case PlatformName:
				b, err := hex.DecodeString(value)
				if err == nil {
					md.PlatformName = string(b)
				}
			case Manufacturer:
				md.Manufacturer = value
			}
		}
	}
	return md, nil
}

type MacOS struct {
}

func (MacOS) Provider() string { return "macOS" }

func (MacOS) GetHostname() (string, error) {
	cmd := exec.Command("hostname")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (MacOS) GetInstanceID() (string, error) {
	return metadata.PlatformUUID, nil
}

func (MacOS) GetInstanceType() (string, error) {
	return metadata.Model, nil
}

func (MacOS) GetRegion() (string, error) {
	return metadata.RegionInfo, nil
}

func (MacOS) GetZone() (string, error) {
	return "", nil
}

func (MacOS) GetPublicIP() (string, error) {
	cmd := exec.Command("curl", "ifconfig.me")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (MacOS) GetPrivateIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// Get the local address used for this connection
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.String(), nil
}
