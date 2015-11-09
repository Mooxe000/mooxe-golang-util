package mxUtil

import (
	mxu "github.com/Mooxe000/mooxe-golang-util/src/util"
)

type MxSliceItf mxu.MxSliceItf
type MxSlice mxu.MxSlice

var (
	// alias
	Pln = mxu.Pln
	Prf = mxu.Prf

	// util
	TypeOf = mxu.TypeOf
	Dd     = mxu.Dd

	// io
	FileStrFromPath = mxu.FileStrFromPath
)

func PkgName() string {
	return "mxUtil"
}
