//go:build darwin

package imds

import "github.com/si3nloong/imds/provider/macos"

func localMachine() InstanceMetadataService {
	return &macos.MacOS{}
}
