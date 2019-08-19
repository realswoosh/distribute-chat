package net

import "reflect"

func byteToInt32(data []byte) int32{
	var n int32

	n |= int32(data[0])
	n |= int32(data[1]) << 8
	n |= int32(data[2]) << 16
	n |= int32(data[3]) << 24

	return n
}


func HasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}
