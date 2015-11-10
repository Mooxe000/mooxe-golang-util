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
	pln("Hello World!!!")
	dd("Hello Golang!!!")

	pkgName := PkgName()
	prf("Hello %s!!!\n", pkgName)

  ms := MxSlice{
    []string{"1", "2", "3", "4", "5"}}.
    Super
  dd(ms)
  p, l := pop(ms)
  dd(p)
  dd(l)
}
