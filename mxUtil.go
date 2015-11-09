package mxUtil

import (
	mxu "github.com/Mooxe000/mooxe-golang-util/src/util"
)

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
  Len = mxu.Len
  At = mxu.At
  Slice = mxu.Slice
  ToStringSlice = mxu.ToStringSlice

  // type
	TypeOf = mxu.TypeOf
)

func PkgName() string {
	return "mxUtil"
}
