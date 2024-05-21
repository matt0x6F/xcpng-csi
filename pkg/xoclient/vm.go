package xoclient

import (
	"errors"
)

// GetVMByName
func (c *xoClient) GetVMByName(name string) ([]*VM, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	// resp should be map[string]interface{} so we will treat it as such

	filters := map[string]string{
		"type":       "VM",
		"name_label": name,
	}

	vms := make([]*VM, 0)

	for _, v := range (*resp).(map[string]interface{}) {
		vm := new(VM)
		if valid := c.Filter(v, filters, vm); valid {
			vms = append(vms, vm)
		}
	}

    return vms, nil
}

// GetVMByUUID
func (c *xoClient) GetVMByUUID(ref VMRef) (*VM, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	// resp should be map[string]interface{} so we will treat it as such

	filters := map[string]string{
		"type": "VM",
		"uuid": string(ref),
	}

	vms := make([]*VM, 0)

	for _, v := range (*resp).(map[string]interface{}) {
		vm := new(VM)
		if valid := c.Filter(v, filters, vm); valid {
			vms = append(vms, vm)
		}
	}

	if len(vms) == 1 {
		return vms[0], nil
	}

	return nil, errors.New("No VM found with this uuid")
}
