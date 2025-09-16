package dkgocr

import (
	"github.com/smartcontractkit/smdkg/dkgocr/dkgocrtypes"
	"github.com/smartcontractkit/smdkg/utils/ocr/plugin"
)

// Create a new empty instance of dkgocrtypes.ResultPackage.
// Used for Unmarshaling via the implementation of the encoding.BinaryUnmarshaler interface.
func NewResultPackage() dkgocrtypes.ResultPackage {
	return &plugin.ResultPackage{}
}
