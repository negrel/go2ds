package elements

import "unsafe"

// Element define any element contained in a data structure.
type Element unsafe.Pointer

var a = 0
var size = unsafe.Sizeof(&a)

// Size return the size of any Element.
func Size() uintptr {
	return size
}
