package util

import (
	"testing"
)

func getMxListIns() (MxList, MxList) {
	var ssml, iiml MxList

	ssmi := MxInterface{[]string{"1", "2", "3", "4", "5"}}
	ssml.Value = ssmi.ToInterfaces()

	iimi := MxInterface{[]int{1, 2, 3, 4, 5}}
	iiml.Value = iimi.ToInterfaces()

	return ssml, iiml
}

func Test_New(t *testing.T) {
	FuncName := "MxLists.New()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	// dd(Pssml.Value)
	// dd(Pssml.list)

	if Pssml.list.Len() != 5 {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

// func Test_Sync(t *testing.T) {
//   FuncName := "MxLists.Sync()"
//   t.Log(FuncName + " ... ok!")
// }

func Test_Len(t *testing.T) {
	FuncName := "MxLists.Len()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	l := Pssml.Len()
	a := len(Pssml.Value)
	b := Pssml.list.Len()

	// dd(v)
	// dd(a)
	// dd(b)

	if a != l || b != l {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_At(t *testing.T) {
	FuncName := "MxLists.At()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	l := Pssml.Len()
	for i := 0; i < l; i++ {
		v := Pssml.At(i).Value
		// dd(v)
		if v != Pssml.Value[i] {
			t.Error(FuncName + " ... failed!")
		}
	}

	t.Log(FuncName + " ... ok!")
}

func Test_Setter(t *testing.T) {
	FuncName := "MxLists.Setter()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	var v interface{} = "100"
	// dd(v)
	Pssml.Setter(2, v)

	// dd(Pssml.Value[2])

	if Pssml.Value[2] != "100" {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_Remove(t *testing.T) {
	FuncName := "MxLists.Remove()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	Pssml.Remove(2).Sync()

	// dd(Pssml.Value)

	if Pssml.Len() != 4 || Pssml.Value[2] != "4" {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_Remove_Sync(t *testing.T) {
	FuncName := "MxLists.Remove() AND MxLists.Sync()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	v := Pssml.Remove(2)

	l := Pssml.Len()
	a := len(Pssml.Value)
	b := Pssml.list.Len()
	// dd(l)
	// dd(a)
	// dd(b)

	if l != b || a-b != 1 {
		t.Error(FuncName + " ... failed!")
	}

	v.Sync()

	ls := Pssml.Len()
	as := len(Pssml.Value)
	// bs := Pssml.list.Len()
	// dd(ls)
	// dd(as)
	// dd(bs)

	if as != ls || ls != 4 || Pssml.Getter(2) != "4" {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_Slice(t *testing.T) {
	FuncName := "MxLists.Slice()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	v := Pssml.Slice(1, 3)

	// dd(v.Value)

	if v.Len() != 3 || v.Value[1] != "3" {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_Chunk_1(t *testing.T) {
	FuncName := "MxLists.Chunk()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	////////////////////////////////
	// Chunk(0/5)
	////////////////////////////////
	mls_0 := Pssml.Chunk(0)
	mls_5 := Pssml.Chunk(5)

	// pln(len(mls_0))
	// pln(len(mls_5))

	if len(mls_0) != 1 || len(mls_5) != 1 {
		t.Error(FuncName + " ... failed!")
	}

	// dd(mls_0[0].Value)
	// dd(mls_5[0].Value)

	if mls_0[0].Len() != len(mls_0[0].Value) || mls_5[0].Len() != 5 {
		t.Error(FuncName + " ... failed!")
	}
	////////////////////////////////

	////////////////////////////////
	// Chunk(1)
	////////////////////////////////
	mls_1 := Pssml.Chunk(1)

	if len(mls_1) != 5 {
		t.Error(FuncName + " ... failed!")
	}

	for _, v := range mls_1 {
		// dd(v.Value)
		if len(v.Value) != 1 {
			t.Error(FuncName + " ... failed!")
		}
	}
	////////////////////////////////

	////////////////////////////////
	// Chunk(2)
	////////////////////////////////
	mls_2 := Pssml.Chunk(2)

	is := []int{2, 2, 1}
	for i, v := range mls_2 {
		// dd(v.Value)
		if v.Len() != is[i] {
			t.Error(FuncName + " ... failed!")
		}
	}
	////////////////////////////////

	////////////////////////////////
	// Chunk(3)
	////////////////////////////////
	mls_3 := Pssml.Chunk(3)

	ii := 0
	for _, v := range mls_3 {
		// dd(v.Value)
		ii = ii + v.Len()
	}

	if ii != 5 {
		t.Error(FuncName + " ... failed!")
	}

	////////////////////////////////
	// Chunk(4)
	////////////////////////////////
	mls_4 := Pssml.Chunk(4)

	if len(mls_4) != 2 {
		t.Error(FuncName + " ... failed!")
	}

	// dd(mls_4[0].Value)
	// dd(mls_4[1].Value)

	if len(mls_4[0].Value) != 4 {
		t.Error(FuncName + " ... failed!")
	}
	if len(mls_4[1].Value) != 1 {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_Chunk_2(t *testing.T) {
	FuncName := "MxLists.Chunk()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	////////////////////////////////
	// Chunk(0, 0, 0)
	////////////////////////////////
	mls_0 := Pssml.Chunk(0, 0, 0)

	if len(mls_0) != 1 {
		t.Error(FuncName + " ... failed!")
	}

	for _, v := range mls_0 {
		// dd(v.Value)
		if len(v.Value) != 5 {
			t.Error(FuncName + " ... failed!")
		}
	}

	////////////////////////////////
	// Chunk(1, 1, 1)
	////////////////////////////////
	mls_1 := Pssml.Chunk(1, 1, 1)

	if len(mls_1) != 4 {
		t.Error(FuncName + " ... failed!")
	}

	for i, v := range mls_1 {
		// dd(v.Value)
		if i <= 2 {
			if len(v.Value) != 1 {
				t.Error(FuncName + " ... failed!")
			}
		} else {
			if i != len(mls_1)-1 || len(v.Value) != 2 {
				t.Error(FuncName + " ... failed!")
			}
		}
	}

	////////////////////////////////
	// TODO Chunk
	////////////////////////////////
	// mls := Pssml.Chunk(1, 3, 8)
	// mls := Pssml.Chunk(3, 0, 1)
	// mls := Pssml.Chunk(0, 2, 5)
	// mls := Pssml.Chunk(2, 1, 2)
	// for _, v := range mls {
	//   dd(v.Value)
	// }

	t.Log(FuncName + " ... ok!")
}

func Test_Concat(t *testing.T) {
	FuncName := "MxLists.Concat()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	Pssml.Concat(Pssml.Chunk(2))

	// dd(Pssml.Value)

	if Pssml.Len() != 5 || Pssml.Value[2] != "3" {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_RemoveList(t *testing.T) {
	FuncName := "MxLists.RemoveList()"

	ssml, _ := getMxListIns()
	Pssml := &ssml
	Pssml.New()

	// l := Pssml.Len()

	// Pssml.RemoveList(-1, -1)
	// Pssml.RemoveList(l, l)
	// Pssml.RemoveList(0, 0)
	// Pssml.RemoveList(l-1, l-1)
	// Pssml.RemoveList(2, 2)

	// Pssml.RemoveList(0, l)
	// Pssml.RemoveList(0, l-1)
	// Pssml.RemoveList(-1, l-1)

	// Pssml.RemoveList(-1, -4)
	// Pssml.RemoveList(4, 10)
	// Pssml.RemoveList(2, -5)
	// Pssml.RemoveList(2, 4)
	// Pssml.RemoveList(2, 3)

	// dd(Pssml.Sync().Value)

	t.Log(FuncName + " ... ok!")
}
