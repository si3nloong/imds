package imds

func GetInstanceID() (string, error) {
	return defaultImds.GetInstanceID()
}

func GetInstanceType() (string, error) {
	return defaultImds.GetInstanceType()
}

func GetRegion() (string, error) {
	return defaultImds.GetRegion()
}

func GetZone() (string, error) {
	return defaultImds.GetZone()
}

func GetPublicIP() (string, error) {
	return defaultImds.GetPublicIP()
}

func GetPrivateIP() (string, error) {
	return defaultImds.GetPrivateIP()
}
