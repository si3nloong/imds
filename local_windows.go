//go:build windows

package imds

import "github.com/si3nloong/imds/provider/windows"

func localMachine() InstanceMetadataService {
	return &windows.Windows{}
}
