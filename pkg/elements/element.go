package elements

import "unsafe"

type Element unsafe.Pointer

var a = 0
var size = unsafe.Sizeof(&a)

func Size() uintptr {
	return size
}
