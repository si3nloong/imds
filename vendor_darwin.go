//go:build darwin

package imds

import "os"

func instanceVendor() (string, error) {
	vendor, err := os.ReadFile("/sys/class/dmi/id/sys_vendor")
	if err != nil {
		return "", err
	}
	return string(vendor), nil
}
