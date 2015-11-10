package main

import (
	. "mxUtil"
)

var (
	dd  = Dd
	pln = Pln
	prf = Prf
  pop = Pop
)

func main() {
	pln("Hello World!!!")
	dd("Hello Golang!!!")

	pkgName := PkgName()
	prf("Hello %s!!!\n", pkgName)

  ms := (MxSlice{MxSlice: []string{"1", "2", "3", "4", "5"}}).MxSlice
  dd(ms)
  p, l := pop(ms)
  dd(p)
  dd(l)
}
