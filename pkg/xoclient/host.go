package xoclient

import "errors"

// GetHostByName
func (c *xoClient) GetHostByName(name string) (*Host, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	// resp should be a map[string]interface{} and will be treated as such

	hosts := make([]*Host, 0)
	filters := map[string]string{
		"type":       "host",
		"name_label": name,
	}

	for _, v := range (*resp).(map[string]interface{}) {
		host := new(Host)
		if valid := c.Filter(v, filters, host); valid {
			hosts = append(hosts, host)
		}
	}

	if len(hosts) == 1 {
		return hosts[0], nil
	}

	return nil, errors.New("Failed finding host")
}

// GetHostByUUID
func (c *xoClient) GetHostByUUID(ref HostRef) (*Host, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	// resp should be a map[string]interface{} and will be treated as such

	hosts := make([]*Host, 0)
	filters := map[string]string{
		"type": "host",
		"uuid": string(ref),
	}

	for _, v := range (*resp).(map[string]interface{}) {
		host := new(Host)
		if valid := c.Filter(v, filters, host); valid {
			hosts = append(hosts, host)
		}
	}

	if len(hosts) == 1 {
		return hosts[0], nil
	}

	return nil, errors.New("Failed finding host")
}
