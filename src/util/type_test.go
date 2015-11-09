package util

import (
	"testing"
)

var (
	typeof = TypeOf
)

func Test_typeof(t *testing.T) {
	if r := typeof("Hello World!!!"); r != "string" {
		t.Error("typeof() ... failed!")
	} else {
		t.Log("typeof() ... ok!")
	}
}
