// AWS
package imds

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gofrs/flock"
)

type AWS struct {
	rw          sync.RWMutex
	token       string
	expiresTime time.Time
}

func (a *AWS) Provider() string {
	return "AWS"
}

type InstanceIdentityDocuments struct {
	AccountID               string    `json:"accountId"`
	Architecture            string    `json:"architecture"`
	AvailabilityZone        string    `json:"availabilityZone"`
	BillingProducts         any       `json:"billingProducts"`
	DevpayProductCodes      any       `json:"devpayProductCodes"`
	MarketplaceProductCodes any       `json:"marketplaceProductCodes"`
	ImageID                 string    `json:"imageId"`
	InstanceID              string    `json:"instanceId"`
	InstanceType            string    `json:"instanceType"`
	KernelID                *string   `json:"kernelId"`
	PendingTime             time.Time `json:"pendingTime"`
	PrivateIP               string    `json:"privateIp"`
	RamdiskID               *string   `json:"ramdiskId"`
	Region                  string    `json:"region"`
	Version                 string    `json:"version"`
}

func (c *AWS) GetInstanceDocument() (InstanceIdentityDocuments, error) {
	body, err := c.curl("/latest/dynamic/instance-identity/document")
	if err != nil {
		return InstanceIdentityDocuments{}, err
	}

	var doc InstanceIdentityDocuments
	if err := json.Unmarshal(body, &doc); err != nil {
		return InstanceIdentityDocuments{}, err
	}
	return doc, nil
}

func (c *AWS) GetHostname() (string, error) {
	b, err := c.curl("/latest/meta-data/hostname")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AWS) GetLocalIPv4() (string, error) {
	b, err := c.curl("/latest/meta-data/local-ipv4")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AWS) GetPublicIPv4() (string, error) {
	b, err := c.curl("/latest/meta-data/public-ipv4")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AWS) GetMAC() (string, error) {
	b, err := c.curl("/latest/meta-data/mac")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AWS) GetInstanceID() (string, error) {
	b, err := c.curl("/latest/meta-data/instance-id")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AWS) GetInstanceType() (string, error) {
	b, err := c.curl("/latest/meta-data/instance-type")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *AWS) renewToken() (string, error) {
	c.rw.RLock()
	// If the token is expired
	if time.Now().Before(c.expiresTime) {
		c.rw.RUnlock()
		return c.token, nil
	}

	c.rw.RUnlock()
	c.rw.Lock()
	defer c.rw.Unlock()
	fileLock := flock.New("/var/lock/go-lock.lock")
	locked, err := fileLock.TryLock()
	if err != nil {
		return "", err
	}
	// If cannot locked, mean other process is renewing the token
	if !locked {
		var ok bool
	wait_for_new_token:
		for {
			// Check the env until it got the new token
			c.token, ok = os.LookupEnv("AWS_IMDS_TOKEN")
			if ok {
				break wait_for_new_token
			}
			<-time.After(time.Second)
		}
		return c.token, nil
	}
	defer fileLock.Unlock()

	// Double check
	req, err := http.NewRequest(http.MethodPut, "http://169.254.169.254/latest/api/token", nil)
	if err != nil {
		return "", err
	}

	expiresIn := 21600
	req.Header.Add("X-aws-ec2-metadata-token-ttl-seconds", strconv.Itoa(expiresIn))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	token := string(b)
	if err := os.Setenv("AWS_IMDS_TOKEN", token); err != nil {
		return "", err
	}
	c.token = token
	c.expiresTime = time.Now().Add(time.Second * time.Duration(expiresIn))
	return token, nil
}

func (c *AWS) curl(path string) ([]byte, error) {
	token, err := c.renewToken()
	if err != nil {
		return nil, fmt.Errorf(`unable to renew token: %w`, err)
	}

	req, err := http.NewRequest(http.MethodGet, "http://169.254.169.254"+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-aws-ec2-metadata-token", token)

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
