//go:build darwin

package imds

import "golang.org/x/sys/unix"

func instanceVendor() (string, error) {
	vendor, err := unix.Sysctl("machdep.cpu.vendor")
	if err == nil {
		return vendor, nil
	}
	return unix.Sysctl("machdep.cpu.brand_string")
}
