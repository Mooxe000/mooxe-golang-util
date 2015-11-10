package util

import (
	"testing"
)

func Test_AutoInterface(t *testing.T) {
  FuncName := "AutoInterface.ToInterfaces()"
  ss := []string{"1", "2", "3", "4", "5"}
  mi := MxInterface{ss}
  Pmi := &mi
  ins := Pmi.ToInterfaces()

  // dd(typeof(ins) == "[]interface {}")
  // dd(IsSlice(ins))

  if IsSlice(ins) && typeof(ins) == "[]interface {}" {
    for i, v := range ins {
      if ss[i] != v {
        t.Error(FuncName + " ... failed!")
      }
    }
  } else {
    t.Error(FuncName + " ... failed!")
  }

  t.Log(FuncName + "... ok!")
}
