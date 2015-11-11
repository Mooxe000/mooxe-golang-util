package util

import (
	"reflect"
)

type MxInterface struct {
	Value interface{}
}

func (mi *MxInterface) TypeOf() string {
	return typeof(mi.Value)
}

func (mi *MxInterface) IsSlice() bool {
	return IsSlice(mi.Value)
}

func (mi *MxInterface) ToInterfaces() []interface{} {
	var r []interface{}
	if mi.IsSlice() { // 基于反射的迭代器
		value := reflect.ValueOf(mi.Value)
		for i := 0; i < value.Len(); i++ {
			v := value.Index(i).Interface()
			r = append(r, v)
		}
	} else {
		return nil
	}
	return r
}
