// Digital Ocean
package imds

type DigitalOcean struct {
}

func (d DigitalOcean) Provider() string {
	return "Digital Ocean"
}
