package alicloud

import (
	"fmt"
	"io"
	"net/http"
)

const (
	Endpoint = "http://100.100.100.200"
)

type AliCloud struct{}

func (a AliCloud) Provider() string {
	return "Alibaba Cloud"
}

func (c *AliCloud) GetHostname() (string, error) {
	b, err := curl("/latest/meta-data/hostname")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) InstanceID() (string, error) {
	b, err := curl("/latest/meta-data/instance-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) InstanceType() (string, error) {
	b, err := curl("/latest/meta-data/instance-type")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) Region() (string, error) {
	b, err := curl("/latest/meta-data/region-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) Zone() (string, error) {
	b, err := curl("/latest/meta-data/zone-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) PublicIP() (string, error) {
	b, err := curl("/latest/meta-data/eipv4")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) PrivateIP() (string, error) {
	b, err := curl("/latest/meta-data/private-ipv4")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetImageID() (string, error) {
	b, err := curl("/latest/meta-data/image-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func curl(url string, onBeforeRequest ...func(*http.Request)) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, Endpoint+url, nil)
	if err != nil {
		return nil, err
	}

	if len(onBeforeRequest) > 0 {
		onBeforeRequest[0](req)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get metadata: %s", res.Status)
	}
	return io.ReadAll(res.Body)
}
