package imds

import (
	"strings"

	"github.com/si3nloong/imds/provider/alicloud"
	"github.com/si3nloong/imds/provider/aws"
	"github.com/si3nloong/imds/provider/azure"
)

var defaultImds InstanceMetadataService

type InstanceMetadataService interface {
	Provider() string
	GetHostname() (string, error)
	GetInstanceID() (string, error)
	GetInstanceType() (string, error)
	GetRegion() (string, error)
	GetZone() (string, error)
	GetPublicIP() (string, error)
	GetPrivateIP() (string, error)
}

func Default() InstanceMetadataService {
	return defaultImds
}

func init() {
	vendor, _ := instanceVendor()
	println(vendor)
	vendor = strings.ToLower(strings.TrimSpace(vendor))
	switch {
	case strings.Contains(vendor, "amazon ec2"):
		defaultImds = aws.New()
	case strings.Contains(vendor, "alibaba cloud"):
		defaultImds = alicloud.New()
	case strings.Contains(vendor, "microsoft corporation"):
		defaultImds = azure.New()
	default:
		defaultImds = localMachine()
	}
}

// func curl(url string, onBeforeRequest ...func(*http.Request)) ([]byte, error) {
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(onBeforeRequest) > 0 {
// 		onBeforeRequest[0](req)
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("failed to get metadata: %s", res.Status)
// 	}
// 	return io.ReadAll(res.Body)
// }
