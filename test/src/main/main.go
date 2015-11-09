package main

import (
	. "github.com/Mooxe000/mooxe-golang-util"
)

var (
	dd  = Dd
	pln = Pln
	prf = Prf
  pop = Pop
)

func main() {
	Pln("Hello World!!!")
	dd("Hello Golang!!!")

	pkgName := PkgName()
	Prf("Hello %s!!!\n", pkgName)

  ms := MxSlice{"1", "2", "3", "4", "5"}
  p, l := pop(ms)
  dd(ms)
  dd(p)
  dd(l)
}
