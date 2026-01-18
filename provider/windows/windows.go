//go:build windows

package windows

import (
	"errors"
	"net"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

type Windows struct{}

func (Windows) Provider() string { return "windows" }

func (Windows) GetInstanceID() (string, error) {
	keys, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\DESCRIPTION\System\BIOS`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer keys.Close()

	productID, _, err := keys.GetStringValue("SystemSKU")
	return productID, err
}

func (Windows) GetInstanceType() (string, error) {
	keys, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\DESCRIPTION\System\BIOS`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer keys.Close()

	productID, _, err := keys.GetStringValue("SystemProductName")
	return productID, err
}

func (Windows) GetRegion() (string, error) {
	return "", errors.New(`no zone info available on windows`)
}

func (Windows) GetZone() (string, error) {
	return "", errors.New(`no zone info available on windows`)
}

func (Windows) GetPublicIP() (string, error) {
	cmd := exec.Command("curl", "ifconfig.me")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (Windows) GetPrivateIP() (string, error) {
	// Get all network interfaces
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// Check if the address is a network IP address
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// Check if it is an IPv4 address and not link-local (169.254.x.x)
			if ipnet.IP.To4() != nil && !ipnet.IP.IsLinkLocalUnicast() {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New(`unable to find machine local ip`)
}

// BaseBoardManufacturer     : Microsoft Corporation
// BaseBoardProduct          : Virtual Machine
// BaseBoardVersion          : Hyper-V UEFI Release v4.1
// BIOSReleaseDate           : 03/08/2024
// BIOSVendor                : Microsoft Corporation
// BIOSVersion               : Hyper-V UEFI Release v4.1
// SystemFamily              : Virtual Machine
// SystemManufacturer        : Microsoft Corporation
// SystemProductName         : Virtual Machine
// SystemSKU                 : None
// SystemVersion             : Hyper-V UEFI Release v4.1

type SystemBiosInfo struct {
	BaseBoardManufacturer string
	BaseBoardProduct      string
	BaseBoardVersion      string
	BIOSReleaseDate       string
	BIOSVendor            string
	BIOSVersion           string
	SystemFamily          string
	SystemManufacturer    string
	SystemProductName     string
	SystemSKU             string
	SystemVersion         string
}

func (Windows) SystemBiosInformation() (SystemBiosInfo, error) {
	keys, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\DESCRIPTION\System\BIOS`, registry.QUERY_VALUE)
	if err != nil {
		return SystemBiosInfo{}, err
	}
	defer keys.Close()

	var o SystemBiosInfo
	valueNames, err := keys.ReadValueNames(-1) // -1 means "read all"
	if err == nil {
		for _, name := range valueNames {
			val, _, _ := keys.GetStringValue(name)
			if val == "" {
				continue
			}
			switch name {
			case "BaseBoardManufacturer":
				o.BaseBoardManufacturer = val
			case "BaseBoardProduct":
				o.BaseBoardProduct = val
			case "BaseBoardVersion":
				o.BaseBoardVersion = val
			case "BIOSReleaseDate":
				o.BIOSReleaseDate = val
			case "BIOSVendor":
				o.BIOSVendor = val
			case "BIOSVersion":
				o.BIOSVersion = val
			case "SystemFamily":
				o.SystemFamily = val
			case "SystemManufacturer":
				o.SystemManufacturer = val
			case "SystemProductName":
				o.SystemProductName = val
			case "SystemSKU":
				o.SystemSKU = val
			case "SystemVersion":
				o.SystemVersion = val
			}
		}
	}
	return o, nil
}

// SystemRoot                : C:\Windows
// BuildBranch               : ge_release
// BuildGUID                 : ffffffff-ffff-ffff-ffff-ffffffffffff
// BuildLab                  : 26100.ge_release.240331-1435
// BuildLabEx                : 26100.1.amd64fre.ge_release.240331-1435
// CompositionEditionID      : Enterprise
// CurrentBuild              : 26200
// CurrentBuildNumber        : 26200
// CurrentType               : Multiprocessor Free
// CurrentVersion            : 6.3
// DisplayVersion            : 25H2
// EditionID                 : Professional
// InstallationType          : Client
// LCUVer                    : 10.0.26100.7462
// ProductName               : Windows 10 Pro
// ReleaseId                 : 2009
// SoftwareType              : System
// PathName                  : C:\Windows
// ProductId                 : 00331-10000-00001-AA295

type SoftwareInfo struct {
	SystemRoot           string
	BuildBranch          string
	BuildGUID            string
	BuildLab             string
	BuildLabEx           string
	CompositionEditionID string
	CurrentBuild         string
	CurrentBuildNumber   string
	CurrentType          string
	CurrentVersion       string
	DisplayVersion       string
	EditionID            string
	InstallationType     string
	LCUVer               string
	ProductName          string
	ReleaseId            string
	SoftwareType         string
	PathName             string
	ProductId            string
}

func (Windows) SoftwareInformation() (SoftwareInfo, error) {
	keys, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return SoftwareInfo{}, err
	}
	defer keys.Close()

	var o SoftwareInfo
	valueNames, err := keys.ReadValueNames(-1) // -1 means "read all"
	if err == nil {
		for _, name := range valueNames {
			val, _, _ := keys.GetStringValue(name)
			if val == "" {
				continue
			}
			switch name {
			case "SystemRoot":
				o.SystemRoot = val
			case "BuildBranch":
				o.BuildBranch = val
			case "BuildGUID":
				o.BuildGUID = val
			case "BuildLab":
				o.BuildLab = val
			case "BuildLabEx":
				o.BuildLabEx = val
			case "CompositionEditionID":
				o.CompositionEditionID = val
			case "CurrentBuild":
				o.CurrentBuild = val
			case "CurrentBuildNumber":
				o.CurrentBuildNumber = val
			case "CurrentType":
				o.CurrentType = val
			case "CurrentVersion":
				o.CurrentVersion = val
			case "DisplayVersion":
				o.DisplayVersion = val
			case "EditionID":
				o.EditionID = val
			case "InstallationType":
				o.InstallationType = val
			case "LCUVer":
				o.LCUVer = val
			case "ProductName":
				o.ProductName = val
			case "ReleaseId":
				o.ReleaseId = val
			case "SoftwareType":
				o.SoftwareType = val
			case "PathName":
				o.PathName = val
			case "ProductId":
				o.ProductId = val
			}
		}
	}
	return o, nil
}
