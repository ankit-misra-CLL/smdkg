package dkgocr

import (
	"github.com/ankit-misra-CLL/smdkg/dkgocr/dkgocrtypes"
	"github.com/ankit-misra-CLL/smdkg/utils/ocr/plugin"
)

// Create a new empty instance of dkgocrtypes.ResultPackage.
// Used for Unmarshaling via the implementation of the encoding.BinaryUnmarshaler interface.
func NewResultPackage() dkgocrtypes.ResultPackage {
	return &plugin.ResultPackage{}
}
