package xoclient

import (
	"github.com/matt0x6f/xcpng-csi/pkg/xolib"
	"github.com/mitchellh/mapstructure"
)

// Filter : filters
func (c *xoClient) Filter(resp interface{}, filters map[string]string, object interface{}) bool {
	switch val := resp.(type) {
	case map[string]interface{}:
		for k, v := range filters {
			if val[k] != v {
				return false
			}
		}

		mapstructure.Decode(val, object)
		return true
	}

	return false
}

// getAllObjects
func (c *xoClient) getAllObjects() (*xolib.MessageResult, error) {
	request := &xolib.MessageRequest{
		Method: "xo.getAllObjects",
	}

	resp, err := c.Call(request)

	return resp, err
}

// GetAll is just like getAllObjects
func (c *xoClient) GetAll() (*xolib.MessageResult, error) {
	return c.getAllObjects()
}
