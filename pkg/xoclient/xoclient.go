package xoclient

/*
This package is used as an interface for xolib, its mostly used for xcpng-csi so feature support will be limited.
however you can still used xolib since it has full capabilities thanks to its generic like interface

Pull requests are welcome to add more commands :)
*/

import (
	"github.com/matt0x6f/xcpng-csi/pkg/xolib"
)

type (
	xoClient struct {
		xolib.Xolib
	}

	// VDIMethods are all methods that create/modify/get/delete a VDI from an SR
	VDIMethods interface {
		CreateVDI(string, int64, SRRef) (*VDIRef, error)
		DeleteVDI(VDIRef) error
		GetVDIByName(string) ([]*VDI, error)
		GetVDIByUUID(VDIRef) (*VDI, error)
	}

	// VBDMethods are all methods that create/modify/get/delete a VBD from an VM
	VBDMethods interface {
		AttachVBD(VDIRef, VMRef) error
		ConnectVBD(VBDRef) error
		DisconnectVBD(VBDRef) error
		DeleteVBD(VBDRef) error
		GetVBDByName(string) ([]*VBD, error)
		GetVBDByUUID(VBDRef) (*VBD, error)
		GetVBDsFromVDI(VDIRef) ([]*VBD, error)
		GetVBDsFromVM(VMRef) ([]*VBD, error)
	}

	// VMMethods are methods used to create/modify/get/delete a VM from a HOST
	VMMethods interface {
		GetVMByName(string) ([]*VM, error)
		GetVMByUUID(VMRef) (*VM, error)
	}

	// SRMethods are methods used to create/modify/get/delete an SR from a HOST
	SRMethods interface {
		GetSRByUUID(SRRef) (*SR, error)
	}

	// HostMethods are used to modify/add/remove/get a Host from Xen Orchestra
	HostMethods interface {
		GetHostByName(string) (*Host, error)
		GetHostByUUID(HostRef) (*Host, error)
	}

	// XOClient is the main interface used to interact with xo client
	XOClient interface {
		VDIMethods
		VBDMethods
		VMMethods
		SRMethods
		HostMethods
		xolib.Xolib
		GetAll() (*xolib.MessageResult, error)
	}
)

// NewClient returns a client
func NewClient(lib xolib.Xolib) XOClient {
	return &xoClient{
		Xolib: lib,
	}
}
