package imds

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const (
	AWSEndpoint      = "http://169.254.169.254"
	AzureEndpoint    = "http://169.254.169.254"
	AliCloudEndpoint = "http://100.100.100.200"
)

var defaultImds InstanceMetadataService

type InstanceMetadataService interface {
	Provider() string
	GetInstanceID() (string, error)
	GetInstanceType() (string, error)
	GetRegion() (string, error)
}

func init() {
	switch runtime.GOOS {
	case "windows":
		// Not always the case because other cloud can also provide windows machine
		defaultImds = &Azure{}
	default:
		data, err := os.ReadFile("/sys/class/dmi/id/sys_vendor")
		if err != nil {
			panic(`missing vendor file`)
		}
		vendor := strings.ToLower(strings.TrimSpace(string(data)))
		switch {
		case strings.Contains(vendor, "amazon"):
			defaultImds = &AWS{}
		case strings.Contains(vendor, "alibaba"):
			defaultImds = &AliCloud{}
		default:
		}
	}
}

func curl(url string, onBeforeRequest ...func(*http.Request)) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if len(onBeforeRequest) > 0 {
		onBeforeRequest[0](req)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get metadata: %s", res.Status)
	}
	return io.ReadAll(res.Body)
}
