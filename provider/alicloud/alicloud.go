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

func (c *AliCloud) GetInstanceID() (string, error) {
	b, err := curl("/latest/meta-data/instance-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetInstanceType() (string, error) {
	b, err := curl("/latest/meta-data/instance-type")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetRegion() (string, error) {
	b, err := curl("/latest/meta-data/region-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetZone() (string, error) {
	b, err := curl("/latest/meta-data/zone-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetPublicIP() (string, error) {
	b, err := curl("/latest/meta-data/eipv4")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AliCloud) GetPrivateIP() (string, error) {
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
