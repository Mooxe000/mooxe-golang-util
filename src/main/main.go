package main

import (
	. "util"
)

var (
	pln    = Pln
	dd     = Dd
	typeof = TypeOf
)

func main() {
	pln("Hello World!!!")
	dd("Hello Golang!!!")
	pln(typeof("Hello"))

  ms := MxSlice{"1", "2", "3", "4", "5"}
  p, l := Pop(ms)
  dd(ms)
  dd(p)
  dd(l)
}
