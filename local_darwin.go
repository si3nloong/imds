//go:build darwin

package imds

import "github.com/si3nloong/imds/provider/os/macos"

func localMachine() InstanceMetadataService {
	return macos.MacOS{}
}
