package main

import (
	. "github.com/Mooxe000/mooxe-golang-util"
)

var (
	dd  = Dd
	pln = Pln
	prf = Prf
)

func main() {
	Pln("Hello World!!!")
  dd("Hello Golang!!!")

	pkgName := PkgName()
	Prf("Hello %s!!!\n", pkgName)
}
