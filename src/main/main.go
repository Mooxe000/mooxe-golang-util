package main

import (
  . "util"
)

var (
  pln = Pln
  dd = Dd
  typeof = TypeOf
)

func main() {
  pln("Hello World!!!")
  dd("Hello Golang!!!")
  pln(typeof("Hello"))
}
