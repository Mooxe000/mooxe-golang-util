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

	// helper
	Dd     = mxu.Dd

	// io
	FileStrFromPath = mxu.FileStrFromPath

  // slice
  Pop = mxu.Pop

  // type
	TypeOf = mxu.TypeOf
)

func PkgName() string {
	return "mxUtil"
}
