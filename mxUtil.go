package mxUtil

import (
  mxu "github.com/Mooxe000/mooxe-golang-util/src/util"
)

var (
  // alias
  Pln = mxu.Pln
  Prf = mxu.Prf

  // util
  TypeOf = mxu.TypeOf
  Dd = mxu.Dd

  // io
  FileStrFromPath = mxu.FileStrFromPath
)

var (
  pln = Pln
)

func PkgName() (string) {
  pkgName := "mxUtil"
  pln(pkgName)
  return pkgName
}
