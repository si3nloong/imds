package azure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

const Endpoint = "http://169.254.169.254"

type resultFormat uint8

const (
	resultFormatText resultFormat = iota
	resultFormatJSON
)

type Azure struct {
}

func New() *Azure {
	http.DefaultClient = &http.Client{
		Transport: &http.Transport{Proxy: nil},
	}
	return &Azure{}
}

func (a Azure) Provider() string {
	return "Azure"
}

func (a Azure) GetHostname() (string, error) {
	cmd := exec.Command("hostname")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

type Metadata struct {
	Compute struct {
		AzEnvironment              string `json:"azEnvironment"`
		CustomData                 string `json:"customData"`
		EvictionPolicy             string `json:"evictionPolicy"`
		IsHostCompatibilityLayerVM string `json:"isHostCompatibilityLayerVm"`
		LicenseType                string `json:"licenseType"`
		Location                   string `json:"location"`
		Name                       string `json:"name"`
		Offer                      string `json:"offer"`
		OsProfile                  struct {
			AdminUsername                 string `json:"adminUsername"`
			ComputerName                  string `json:"computerName"`
			DisablePasswordAuthentication string `json:"disablePasswordAuthentication"`
		} `json:"osProfile"`
		OsType           string `json:"osType"`
		PlacementGroupID string `json:"placementGroupId"`
		Plan             struct {
			Name      string `json:"name"`
			Product   string `json:"product"`
			Publisher string `json:"publisher"`
		} `json:"plan"`
		PlatformFaultDomain  string        `json:"platformFaultDomain"`
		PlatformUpdateDomain string        `json:"platformUpdateDomain"`
		Priority             string        `json:"priority"`
		Provider             string        `json:"provider"`
		PublicKeys           []interface{} `json:"publicKeys"`
		Publisher            string        `json:"publisher"`
		ResourceGroupName    string        `json:"resourceGroupName"`
		ResourceID           string        `json:"resourceId"`
		SecurityProfile      struct {
			SecureBootEnabled string `json:"secureBootEnabled"`
			VirtualTpmEnabled string `json:"virtualTpmEnabled"`
		} `json:"securityProfile"`
		Sku            string `json:"sku"`
		StorageProfile struct {
			DataDisks      []interface{} `json:"dataDisks"`
			ImageReference struct {
				ID        string `json:"id"`
				Offer     string `json:"offer"`
				Publisher string `json:"publisher"`
				Sku       string `json:"sku"`
				Version   string `json:"version"`
			} `json:"imageReference"`
			OsDisk struct {
				Caching          string `json:"caching"`
				CreateOption     string `json:"createOption"`
				DiffDiskSettings struct {
					Option string `json:"option"`
				} `json:"diffDiskSettings"`
				DiskSizeGB         string `json:"diskSizeGB"`
				EncryptionSettings struct {
					Enabled string `json:"enabled"`
				} `json:"encryptionSettings"`
				Image struct {
					URI string `json:"uri"`
				} `json:"image"`
				ManagedDisk struct {
					ID                 string `json:"id"`
					StorageAccountType string `json:"storageAccountType"`
				} `json:"managedDisk"`
				Name   string `json:"name"`
				OsType string `json:"osType"`
				Vhd    struct {
					URI string `json:"uri"`
				} `json:"vhd"`
				WriteAcceleratorEnabled string `json:"writeAcceleratorEnabled"`
			} `json:"osDisk"`
			ResourceDisk struct {
				Size string `json:"size"`
			} `json:"resourceDisk"`
		} `json:"storageProfile"`
		SubscriptionID string `json:"subscriptionId"`
		Tags           string `json:"tags"`
		TagsList       []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"tagsList"`
		UserData       string `json:"userData"`
		Version        string `json:"version"`
		VMID           string `json:"vmId"`
		VMScaleSetName string `json:"vmScaleSetName"`
		VMSize         string `json:"vmSize"`
		Zone           string `json:"zone"`
	} `json:"compute"`
	Network struct {
		Interface []struct {
			Ipv4 struct {
				IPAddress []struct {
					PrivateIPAddress string `json:"privateIpAddress"`
					PublicIPAddress  string `json:"publicIpAddress"`
				} `json:"ipAddress"`
				Subnet []struct {
					Address string `json:"address"`
					Prefix  string `json:"prefix"`
				} `json:"subnet"`
			} `json:"ipv4"`
			Ipv6 struct {
				IPAddress []interface{} `json:"ipAddress"`
			} `json:"ipv6"`
			MacAddress string `json:"macAddress"`
		} `json:"interface"`
	} `json:"network"`
}

func (a *Azure) GetMetadata() (Metadata, error) {
	b, err := a.curl("/metadata/instance", resultFormatJSON)
	if err != nil {
		return Metadata{}, err
	}
	var o Metadata
	if err = json.Unmarshal(b, &o); err != nil {
		return Metadata{}, err
	}
	return o, nil
}

func (a *Azure) InstanceID() (string, error) {
	b, err := a.curl("/metadata/instance/compute/vmId", resultFormatText)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Azure) InstanceType() (string, error) {
	b, err := c.curl("/metadata/instance/compute/vmSize", resultFormatText)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (a *Azure) GetInstanceName() (string, error) {
	b, err := a.curl("/metadata/instance/compute/name", resultFormatText)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (a *Azure) Region() (string, error) {
	b, err := a.curl("/metadata/instance/compute/location", resultFormatText)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (a *Azure) Zone() (string, error) {
	b, err := a.curl("/metadata/instance/compute/zone", resultFormatText)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (a *Azure) PublicIP() (string, error) {
	b, err := a.curl("/metadata/loadbalancer", resultFormatText)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (a *Azure) PrivateIP() (string, error) {
	b, err := a.curl("/metadata/instance/network/interface/0/ipv4/ipAddress/0/privateIpAddress", resultFormatText)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Azure) curl(path string, format resultFormat) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, Endpoint+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Metadata", "True")

	q := req.URL.Query()
	switch format {
	case resultFormatJSON:
		q.Add("format", "json")
	default:
		q.Add("format", "text")
	}
	q.Add("api-version", "2021-02-01")
	req.URL.RawQuery = q.Encode()

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
