package imds

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

func PublicIP() (string, error) {
	return defaultImds.GetPublicIP()
}
