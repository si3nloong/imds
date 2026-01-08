package imds

func Hostname() (string, error) {
	return defaultImds.GetHostname()
}

func InstanceID() (string, error) {
	return defaultImds.GetInstanceID()
}

func InstanceType() (string, error) {
	return defaultImds.GetInstanceType()
}

func Region() (string, error) {
	return defaultImds.GetRegion()
}

func Zone() (string, error) {
	return defaultImds.GetZone()
}

func PublicIPv4() (string, error) {
	return defaultImds.GetPublicIP()
}

func PrivateIPv4() (string, error) {
	return defaultImds.GetPrivateIP()
}

// func PublicIPv6() (string, error) {
// 	return defaultImds.GetPublicIP()
// }

// func PrivateIPv6() (string, error) {
// 	return defaultImds.GetPublicIP()
// }
