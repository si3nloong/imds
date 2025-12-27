package imds

import (
	"encoding/json"
	"net/http"
)

type Azure struct {
}

func (a *Azure) Provider() string {
	return "Azure"
}

func NewAzure() *Azure {
	return &Azure{}
}

func (a *Azure) GetMetadata() (json.RawMessage, error) {
	b, err := a.curl("/metadata/instance")
	if err != nil {
		return nil, err
	}
	return json.RawMessage(b), nil
}

func (a *Azure) GetInstanceID() (string, error) {
	b, err := a.curl("/meta-data/instance/compute/vmId")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Azure) GetInstanceType() (string, error) {
	b, err := c.curl("/meta-data/instance/compute/vmSize")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (a *Azure) GetInstanceName() (string, error) {
	b, err := a.curl("/meta-data/instance/compute/name")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (a *Azure) GetRegion() (string, error) {
	b, err := a.curl("/meta-data/instance/compute/location")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Azure) curl(path string) ([]byte, error) {
	return curl(AzureEndpoint+path, func(r *http.Request) {
		r.Header.Add("Metadata", "True")
	})
}
