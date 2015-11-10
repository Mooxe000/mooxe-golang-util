package main

import (
	. "util"
)

var (
	pln    = Pln
	prf    = Prf
	dd     = Dd
	typeof = TypeOf
)

func print() {
	pln("Hello World!!!")
	dd("Hello Golang!!!")
	pln(typeof("Hello"))
}

func oldSlice() {
	ms := MxSlice{"1", "2", "3", "4", "5"}
	wms := WMxSliceItf{ms}
	p, l := wms.Pop()
	dd(ms)
	dd(p)
	dd(l)
}

func main() {
	ss := []string{"1", "2", "3", "4", "5"}
	ati := AnyToInterface{SS: []string{ss}}
	mc := MxContainer{ati.get("ss")}
	dd(mc)
}
