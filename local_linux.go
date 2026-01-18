//go:build linux

package imds

import "github.com/si3nloong/imds/provider/linux"

func localMachine() InstanceMetadataService {
	return &linux.Linux{}
}
