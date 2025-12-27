package imds

type AliCloud struct {
}

func (a *AliCloud) Provider() string {
	return "AliCloud"
}

func (c *AliCloud) GetInstanceID() (string, error) {
	b, err := curl(AliCloudEndpoint + "/latest/meta-data/instance-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetInstanceType() (string, error) {
	b, err := curl(AliCloudEndpoint + "/latest/meta-data/instance-type")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetRegion() (string, error) {
	b, err := curl(AliCloudEndpoint + "/latest/meta-data/region-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetHostname() (string, error) {
	b, err := curl(AliCloudEndpoint + "/latest/meta-data/hostname")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetPublicIP() (string, error) {
	b, err := curl(AliCloudEndpoint + "/latest/meta-data/eipv4")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetPrivateIP() (string, error) {
	b, err := curl(AliCloudEndpoint + "/latest/meta-data/private-ipv4")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetImageID() (string, error) {
	b, err := curl(AliCloudEndpoint + "/latest/meta-data/image-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
