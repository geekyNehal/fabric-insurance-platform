//go:build !go1.13
// +build !go1.13

package encoder

import "unsafe"

//go:linkname MapIterValue reflect.mapitervalue
func MapIterValue(it *mapIter) unsafe.Pointer
