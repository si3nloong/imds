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
	InstanceID() (string, error)
	InstanceType() (string, error)
	Region() (string, error)
	Zone() (string, error)
	PublicIP() (string, error)
	PrivateIP() (string, error)
}

func Default() InstanceMetadataService {
	return defaultImds
}

func init() {
	vendor, err := instanceVendor()
	if err == nil {
		vendor = strings.ToLower(strings.TrimSpace(vendor))
		switch {
		case strings.Contains(vendor, "amazon ec2"):
			defaultImds = &aws.AWS{}
		case strings.Contains(vendor, "alibaba cloud"):
			defaultImds = &alicloud.AliCloud{}
		case strings.Contains(vendor, "microsoft corporation"):
			defaultImds = azure.New()
		default:
			defaultImds = localMachine()
		}
	} else {
		defaultImds = localMachine()
	}
}
