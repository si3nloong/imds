//go:build windows

package imds

import (
	"golang.org/x/sys/windows/registry"
)

func instanceVendor() (string, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\DESCRIPTION\System\BIOS`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer key.Close()

	// "SystemManufacturer" is the direct equivalent of "sys_vendor"
	vendor, _, err := key.GetStringValue("SystemManufacturer")
	if err != nil {
		return "", err
	}
	return vendor, nil
}
