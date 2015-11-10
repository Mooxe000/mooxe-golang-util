package util

import (
  "container/list"
)

// TODO rename MxList TO MxList
type MxList struct {
  Value []interface{}
  list *list.List
  item interface{}
}

func (ml *MxList) New() *MxList {
  l := list.New()
  for _, v := range ml.Value {
    l.PushBack(v)
  }
  ml.list = l
  return ml
}

func (ml *MxList) Sync() *MxList {
  var r []interface{}
  l := ml.list
  for iter := l.Front();iter != nil ;iter = iter.Next() {
    r = append(r, iter.Value)
  }
  ml.Value = r
  return ml
}

func (ml *MxList) Len() int {
  l := ml.list
  return l.Len()
}

func (ml *MxList) At(i int) *list.Element {
  l := ml.list
  ll := ml.Len()
  var r *list.Element
  if i < 0 || i >= ll {
    return nil
  } else {
    start := l.Front()
    switch i {
    case 0 :
      r = start
    case ll - 1:
      r = l.Back()
    default:
      t := 0
      for e := start ; e != nil; e = e.Next() {
        if t == i {
          r = e
          break
        } else {
          t++
        }
      }
    }
  }
  return r
}

func (ml *MxList) Getter(i int) interface{} {
  return ml.At(i).Value
}

func (ml *MxList) Setter(i int, v interface{}) *MxList {
  ml.At(i).Value = v
  return ml
}

func (ml *MxList) Remove(i int) *MxList {
  l := ml.list
  item := ml.At(i)
  l.Remove(item)
  return ml
}

// Slice
func (ml *MxList) Slice(s int, e int) *MxList {
  var nml MxList
  v := ml.Value
  nml.Value = v[s:e]
  Pml := &nml
  Pml.New()
  return &nml
}

// (Split)
// (Concat)
// (Insert)

// get last
func (ml *MxList) Pop() interface{} {
  leng := ml.Len()-1
  p := ml.Getter(leng)
  ml.Remove(leng)
  return p
}

// add last
func (ml *MxList) Push(i interface{}) *MxList {
  l := ml.list
  l.PushBack(i)
  return ml
}

// get first
func (ml *MxList) Shift() interface{} {
  p := ml.Getter(0)
  ml.Remove(0)
  return p
}

// add first
func (ml *MxList) Unshift(i interface{}) *MxList {
  l := ml.list
  l.PushFront(i)
  return ml
}

// (IndexOf)
