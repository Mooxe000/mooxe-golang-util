package main

import (
  . "github.com/Mooxe000/mooxe-golang-util"
)

var (
  dd = Dd
  pln = Pln
  prf = Prf
)

func main() {
  Pln("Hello World!!!")
  Prf("Hello %s!!!\n", PkgName())
}
