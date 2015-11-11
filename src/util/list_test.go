package util

import (
	"testing"
)

func Test_MxList(t *testing.T) {
  FuncName := "MxList"

  var ml MxList
  mi := MxInterface{[]string{"1", "2", "3", "4", "5"}}
	ml.Value = mi.ToInterfaces()
  Pml := &ml
  Pml.New()
  // Pml.Sync()


  // Push
  // var a interface{}
  // a = "6"
  // Pml.Push(a).Sync()

  for i := 0; i <= 4; i++ {
    // dd(Pml.Getter(i))
    // dd(IntToString(i+1))
    if Pml.Getter(i) != IntToString(i+1) {
      t.Error(FuncName + " ... failed!")
    }
  }

  // dd(Pml.Slice(1, 3).Value)
  // dd(Pml.Value)

  Pml.Split(4, 2)

  // Pml.Remove(4).Push("5").Sync()
  // .Push("5").Sync()
  // a := Pml.Pop()
  // dd(a)
  // dd(Pml.Sync().Value)

  // Pml.Remove(0).Unshift("1").Sync()

  // Pml.Setter(3, "100").Sync()

  // dd(Pml.Value)

	t.Log(FuncName + " ... ok!")
}
