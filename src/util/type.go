package util

import (
	"reflect"
)

func TypeOf(obj interface{}) string {
	result := reflect.TypeOf(obj).String()
	return (result)
}
