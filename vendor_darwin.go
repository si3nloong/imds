//go:build darwin

package imds

import "golang.org/x/sys/unix"

func instanceVendor() (string, error) {
	return unix.Sysctl("machdep.cpu.vendor")
}
