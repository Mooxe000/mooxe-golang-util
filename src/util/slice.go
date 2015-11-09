package util

import (
  "errors"
)

type MxSliceItf interface {
  Len() int
  At(int) interface{}
  Slice(int, int) (MxSlice, error)
}

func Pop(a MxSliceItf) (interface{}, MxSlice) {
  x := a.At(a.Len()-1)
  v, _ := a.Slice(0, a.Len()-1)
  return x, v
}

type MxSlice []string

func (ms MxSlice) Len() int {
  return len(ms)
}

func (ms MxSlice) At(n int) interface{} {
  var r interface{}
  l := len(ms)
  if l!=0 && n>=0 && n<=l {
    r = ms[n]
  }
  return r
}

func (ms MxSlice) Slice(start int, end int) (MxSlice, error) {
  l := len(ms)
  if l!=0 && start>=0 && end<=l {
    return ms[start:end], nil
  } else {
    err := errors.New("Input parameter is incorrect.")
    return nil, err
  }
}

func (ms MxSlice) ToStringSlice() []string {
  a := make([]string, ms.Len())
  for i, v := range ms {
    a[i] = v
  }
  return a
}

// func Push(a []interface{}, x interface{}) ([]interface{}) {
//   a = append(a, x)
//   return a
// }
//
// func Shift(a []interface{}) (interface{}, []interface{}) {
//   x, a := a[0], a[1:]
//   return x, a
// }
//
// func Unshift(a []interface{}, x interface{}) ([]interface{}){
//   r := make([]interface{}, 1)
//   r[0] = x
//   r = append(r, a...)
//   return r
// }
