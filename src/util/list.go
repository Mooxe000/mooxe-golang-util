package util

import (
	"container/list"
	// "sort"
	// "reflect"
)

type MxList struct {
	Value []interface{}
	list  *list.List
}

func (ml *MxList) New() *MxList {
	l := list.New() // TODO ml.list.init()
	for _, v := range ml.Value {
		l.PushBack(v)
	}
	ml.list = l
	return ml
}

// (Sync) TODO arithmetic wait optimize
func (ml *MxList) Sync() *MxList {
	var r []interface{}
	l := ml.list
	for iter := l.Front(); iter != nil; iter = iter.Next() {
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
		case 0:
			r = start
		case ll - 1:
			r = l.Back()
		default:
			t := 0
			for e := start; e != nil; e = e.Next() {
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
	return ml.Value[i]
}

func (ml *MxList) Setter(i int, v interface{}) *MxList {
	ml.Value[i] = v
	return ml
}

// Remove
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
	nml.Value = v[s : e+1]
	Pml := &nml
	Pml.New()
	return &nml
}

// TODO Split To Chunk
func (ml *MxList) Chunk(ii ...int) []MxList {
	lii := len(ii)

	var mls []MxList

	if lii <= 0 {
		mls = append(mls, *ml)
		return mls
	}

	f := ii[0]
	if f < 0 || f >= ml.Len() {

		mls = append(mls, *ml)
		return mls

	} else {

		s, e := 0, f
		// dd(s)
		// dd(e)

		for i := 1; ; i++ {
      if s != e {
  			var nml MxList
  			nml.Value = ml.Slice(s, e-1).Value
  			nml.New()
  			mls = append(mls, nml)
      }

			// start point
			if e >= ml.Len() {
				break
			} else {
				s = e
			}

			// end point
			// -- step n
			var n int
			if lii == 1 { // one repeat

        if f == 0 {

          mls = append(mls, *ml)
      		return mls

        } else{
				  n = f
        }
			} else { // one by one
				if i < lii { // item left
					if ii[i] == 0 {
						continue
					} else {
						n = ii[i]
					}
				} else { // item empty
					n = ml.Len() - s
				}
			}

      if s+n >= ml.Len() {
        n = ml.Len() - s
      }

			// prf("s: %d\n", s)
			// prf("e: %d\n", e)
			// prf("n: %d\n", n)

			e = s + n
		}
	}
	return mls
}

// Concat
func (ml *MxList) Concat(mls []MxList) *MxList {
	var nml MxList
	Pnml := &nml
	for _, v := range mls {
		if Pnml.list == nil {
			Pnml.list = v.list
		} else {
			Pnml.list.PushBackList(v.list)
		}
	}
	ml.list = Pnml.list
	return ml
}

// Remove
// RemoveList
// func (ml *MxList) RemoveList(s int, e int) *MxList {
// 	if s == e {
// 		return ml
// 	}
//   ii := []int{s, e}
//   sort.Ints(ii)
//   dd(ii)
// 	l := ml.Len()
// 	if s <= 0 {
// 		ml.list = ml.Slice(e, l).list
// 	} else if e >= l {
// 		ml.list = ml.Slice(0, s).list
// 	} else {
// 		mls := ml.Slice(s, e) // []MxList
// 		var nmls []MxList
// 		nmls = append(nmls, mls[0])
// 		nmls = append(nmls, mls[len(mls)-1])
// 		ml.Concat(nmls)
// 	}
// 	return ml
// }

// Insert
// InsertList
// Without

// get last
func (ml *MxList) Pop() interface{} {
	leng := ml.Len() - 1
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
// (Flatten)
// (Sort)

////////////////////////////////////////////////////
////////////////////////////////////////////////////
////////////////////////////////////////////////////
// func (ml *MxList) RemoveListReflect(s int, e int) *MxList {
//   if s == e {
//     return ml
//   }
//   if s > e {
//     t := e
//     s = t
//     e = s
//   }
//   l := ml.Len()
//   if s <= 0 {
//     ml.list = ml.Slice(e, l).list
//   } else if e >= l {
//     ml.list = ml.Slice(0, s).list
//   } else {
//     hb := reflect.ValueOf(ml.At(s-1)).Elem()
//     ff := reflect.ValueOf(ml.At(e+1)).Elem()
//
//     // 0 - Next 1 - Prev
//     // 2 - List 3 - Value
//     // dd(hb.Type().Field(n).Name)
//
//     // dd(hb.Field(1).CanSet())
//
//     // dd(hb.Field(3).Interface())
//     // var v interface{} = "-------------"
//     // hb.Field(3).Set(reflect.ValueOf(v))
//     // ff.Field(3).Set(reflect.ValueOf(v))
//
//     hbn := hb.Field(0)
//     ffp := ff.Field(1)
//
//     dd(hbn.CanSet())
//     dd(ffp.CanSet())
//
//     Phb := ml.At(s).Prev()
//     Pff := ml.At(e).Next()
//
//     pln(reflect.ValueOf(Phb).Type())
//     pln(reflect.ValueOf(Pff).Type())
//
//     pln(hbn.Type())
//     pln(ffp.Type())
//
//     // hbn.Set(reflect.ValueOf(&Phb))
//     // ffp.Set(reflect.ValueOf(&Pff))
//   }
//
//   return ml
// }
