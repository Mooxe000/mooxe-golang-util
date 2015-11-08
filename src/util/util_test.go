package util

import (
  "testing"
)

var (
  dd = Dd
  typeof = TypeOf
)

func Test_typeof(t *testing.T) {
  if r := typeof("Hello World!!!"); r != "string" {
    t.Error("typeof() ... failed!")
  } else {
    t.Log("typeof() ... ok!")
  }
}

func Test_dd(t *testing.T) {
  if r, e := dd("Hello World!!!"); r != 17 || e != nil {
    t.Error("dd() ... failed!")
  } else {
    t.Log("dd() ... ok!")
  }
}
