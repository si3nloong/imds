package imds

func InstanceID() (string, error) {
	return defaultImds.InstanceID()
}

func InstanceType() (string, error) {
	return defaultImds.InstanceType()
}

func Region() (string, error) {
	return defaultImds.Region()
}

func Zone() (string, error) {
	return defaultImds.Zone()
}

func PublicIPv4() (string, error) {
	return defaultImds.PublicIP()
}

func PrivateIPv4() (string, error) {
	return defaultImds.PrivateIP()
}
