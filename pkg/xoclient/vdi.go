package xoclient

import (
	"errors"

	"github.com/matt0x6f/xcpng-csi/pkg/xolib"
)

// GetVDIByUUID : ref VDIRef (string)
func (c *xoClient) GetVDIByUUID(ref VDIRef) (*VDI, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	// resp should be a map[string]interface{} so we will treat it a such

	vdis := make([]*VDI, 0)

	filters := map[string]string{
		"type": "VDI",
		"uuid": string(ref),
	}

	for _, v := range (*resp).(map[string]interface{}) {
		vdi := new(VDI)
		if valid := c.Filter(v, filters, vdi); valid {
			vdis = append(vdis, vdi)
		}
	}

	if len(vdis) == 1 {
		return vdis[0], nil
	}

	return nil, errors.New("VDI Not found")
}

// GetVDIByName : name string (pv name)
func (c *xoClient) GetVDIByName(name string) ([]*VDI, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	// resp should be a map[string]interface{} so we will treat it a such

	vdis := make([]*VDI, 0)

	filters := map[string]string{
		"type":       "VDI",
		"name_label": name,
	}

	for _, v := range (*resp).(map[string]interface{}) {
		vdi := new(VDI)
		if valid := c.Filter(v, filters, vdi); valid {
			vdis = append(vdis, vdi)
		}
	}

	return vdis, nil
}

// CreateVDI : Name string, Size int64, SR SRRef
func (c *xoClient) CreateVDI(name string, size int64, sr SRRef) (*VDIRef, error) {
	params := xolib.Params{
		"name": name,
		"size": size,
		"sr":   sr,
	}

	request := xolib.MessageRequest{
		Method: "disk.create",
		Params: &params,
	}

	resp, err := c.Call(&request)
	if err != nil {
		return nil, err
	}

	/*
		Resp here should be a VDIRef so we will treat it as such
	*/

	ref := VDIRef((*resp).(string))

	return &ref, nil

}

// DeleteVDI : deletes VDI
func (c *xoClient) DeleteVDI(ref VDIRef) error {
	params := &xolib.Params{
		"id": string(ref),
	}

	request := &xolib.MessageRequest{
		Method: "vdi.delete",
		Params: params,
	}

	_, err := c.Call(request)
	if err != nil {
		return err
	}

	return nil
}
