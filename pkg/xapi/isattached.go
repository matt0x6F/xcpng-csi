package xapi

import (
	"github.com/matt0x6f/xcpng-csi/pkg/xoclient"
)

func (c *xClient) IsAttached(volID, nodeID string) (bool, error) {
	vdiRef := xoclient.VDIRef(volID)
	vm, err := c.GetVMFromK8sNode(nodeID)
	if err != nil {
		return false, err
	}

	vbds, err := c.GetVBDsFromVM(vm.UUID)
	if err != nil {
		return false, err
	}

	for _, vbd := range vbds {
		if vbd.VDI == vdiRef {
			return true, nil
		}
	}

	return false, nil
}
