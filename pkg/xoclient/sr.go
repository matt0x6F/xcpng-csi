package xoclient

import "errors"

// GetSRByUUID
func (c *xoClient) GetSRByUUID(ref SRRef) (*SR, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	// resp is map[string]interface{} and will be treated as such

	filters := map[string]string{
		"type": "SR",
		"uuid": string(ref),
	}

	srs := make([]*SR, 0)

	for _, v := range (*resp).(map[string]interface{}) {
		sr := new(SR)
		if valid := c.Filter(v, filters, sr); valid {
			srs = append(srs, sr)
		}
	}

	if len(srs) == 1 {
		return srs[0], nil
	}

	return nil, errors.New("Failed at finding SR")
}
