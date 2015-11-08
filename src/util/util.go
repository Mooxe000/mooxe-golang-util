package util

import (
	"github.com/kr/pretty"
	"reflect"
)

func TypeOf(obj interface{}) string {
	result := reflect.TypeOf(obj).String()
	return (result)
}

func Dd(obj interface{}) (int, error) {
	r, e := Prf("%# v\n", pretty.Formatter(obj))
	return r, e
}