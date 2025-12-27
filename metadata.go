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
