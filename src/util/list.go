package util

import (
	"container/list"
	"sort"
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

// TODO params int TO uint avoid border check

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
	if i < 0 || i >= ml.Len() {
		return ml
	}
	l := ml.list
	item := ml.At(i)
	l.Remove(item)
	return ml
}

func iCheck(i int, l int) int {
	var r int
	if i < 0 {
		r = 0
	} else if i >= l {
		r = l - 1
	} else {
		r = i
	}
	return r
}

func seCheck(s int, e int, l int) (int, int) {
	ii := []int{s, e}
	sort.Ints(ii)
	// dd(ii)
	a, b := ii[0], ii[1]
	if a < 0 && b < 0 {
		a, b = 0, 0
	} else if a >= l && b >= l {
		a, b = l-1, l-1
	} else if a < 0 && b >= l {
		a = 0
		b = l - 1
	} else if a < 0 && b >= 0 {
		a = 0
	} else if a >= 0 && b >= l {
		b = l - 1
	}
	return a, b
}

// Slice
func (ml *MxList) Slice(s int, e int) *MxList {
	l := ml.Len()
	a, b := seCheck(s, e, l)

	var nml MxList
	v := ml.Value

	if a == b {
		var is []interface{}
		is = append(is, v[s])
		nml.Value = is

	} else {
		nml.Value = v[a : b+1]
	}

	Pml := &nml
	Pml.New()
	return Pml
}

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

				} else {
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

// RemoveList
func (ml *MxList) RemoveList(s int, e int) *MxList {
	l := ml.Len()
	a, b := seCheck(s, e, l)

	if a == b {
		return ml.Remove(s)

	} else if a == 0 && b == l-1 {
		ml.list.Init()
		return ml

	} else if a == 0 && b < l-1 {
		ml.list = ml.Slice(b+1, l-1).list

	} else if a > 0 && b == l-1 {
		ml.list = ml.Slice(0, a-1).list

	} else { // a > 0 && b < l - 1

		var mls []MxList

		mla := ml.Slice(0, a-1)
		mlb := ml.Slice(b+1, l-1)

		// dd(mla.Value)
		// dd(mlb.Value)

		mls = append(mls, *mla)
		mls = append(mls, *mlb)

		ml.list = ml.Concat(mls).list
	}

	return ml
}

// InsertList
func (ml *MxList) InsertList(i int, nml *MxList) *MxList {
	l := ml.Len()
	i = iCheck(i, l)

	Tml := typeof(ml.Value[0])
	Tnml := typeof(nml.Value[0])
	if Tml != Tnml {
		return ml
	}

	var mls []MxList
	if i == 0 {

		mls = append(mls, *nml)
		mls = append(mls, *ml)
		ml.Concat(mls)

	} else if i == l-1 {

		mls = append(mls, *ml)
		mls = append(mls, *nml)
		ml.Concat(mls)

	} else {

		a := ml.Slice(0, i-1)
		b := ml.Slice(i, l-1)
		mls = append(mls, *a)
		mls = append(mls, *nml)
		mls = append(mls, *b)
		ml.Concat(mls)

	}
	return ml
}

// Insert
func (ml *MxList) Insert(i int, ii interface{}) *MxList {
	l := ml.Len()
	i = iCheck(i, l)

	Tml := typeof(ml.Value[0])
	Tii := typeof(ii)
	if Tml != Tii {
		return ml
	}

	var v []interface{}
	v = append(v, ii)
	var nml MxList
	nml.Value = v
	Pnml := &nml
	Pnml.New()

	ml.InsertList(i, Pnml).Sync()

	return ml
}

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

// (Without)
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
