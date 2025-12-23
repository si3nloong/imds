package imds

import (
	"os"
	"runtime"
	"strings"
)

type InstanceMetadataService interface {
	Provider() string
}

func init() {
	switch runtime.GOOS {
	case "windows":
	default:
		data, err := os.ReadFile("/sys/class/dmi/id/sys_vendor")
		if err != nil {
		}
		vendor := strings.TrimSpace(string(data))
		switch vendor {
		default:
		}
	}
}
