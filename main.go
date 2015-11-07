package main

import (
  "./src"
)

var (
  pln = util.Pln
  dd = util.Dd
  typeof = util.TypeOf
)

func main() {
  pln("Hello World!!!")
  dd("Hello Golang!!!")
  pln(typeof("Hello"))
}
