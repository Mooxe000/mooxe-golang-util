package main

import (
  "../../"
)

var (
  pln = mxUtil.Pln
  dd = mxUtil.Dd
  typeof = mxUtil.TypeOf
)

func main() {
  pln("Hello World!!!")
  dd("Hello Golang!!!")
  pln(typeof("Hello"))
}
