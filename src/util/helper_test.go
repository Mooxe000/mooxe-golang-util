package util

import (
	"testing"
)

var (
	dd = Dd
)

func Test_dd(t *testing.T) {
	if r, e := dd("Hello World!!!"); r != 17 || e != nil {
		t.Error("dd() ... failed!")
	} else {
		t.Log("dd() ... ok!")
	}
}
