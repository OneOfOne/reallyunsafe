package reallyunsafe

import (
	"reflect"
	"unsafe"
)

// NativeBinary gives raw memory access to a []byte, there are *no* bound checking.
// This is extremely unsafe, if misused it could poison your pets and burn your house.
// You've been warned, this is NOT safe.
type NativeBinary uintptr

// NewNativeBinary returns a new NativeBinary with the given slice pointer.
// Again, there are *zero* bounds checking, it is the caller's responsibility to check for safety.
func NewNativeBinary(p *[]byte) NativeBinary {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(p))
	return NativeBinary(sh.Data)
}

// NewStringNativeBinary the same as NewNativeBinary but uses a string pointer for input.
// Again, there are *zero* bounds checking, it is the caller's responsibility to check for safety.
func NewStringNativeBinary(s *string) NativeBinary {
	sh := (*reflect.StringHeader)(unsafe.Pointer(s))
	return NativeBinary(sh.Data)
}

func (n NativeBinary) Uint8(i int) uint8 {
	return *(*uint8)(unsafe.Pointer(uintptr(n) + uintptr(i)))
}

func (n NativeBinary) Uint16(i int) uint16 {
	return *(*uint16)(unsafe.Pointer(uintptr(n) + uintptr(i)))
}

func (n NativeBinary) Uint32(i int) uint32 {
	return *(*uint32)(unsafe.Pointer(uintptr(n) + uintptr(i)))
}

func (n NativeBinary) Uint64(i int) uint64 {
	return *(*uint64)(unsafe.Pointer(uintptr(n) + uintptr(i)))
}

// Slice returns a subslice, appending to it would result in the end of the world or absolutely nothing.
func (n NativeBinary) Slice(i, ln int) []byte {
	sh := &reflect.SliceHeader{
		Data: uintptr(n) + uintptr(i),
		Len:  ln,
		Cap:  ln,
	}
	return *(*[]byte)(unsafe.Pointer(sh))
}

func (n NativeBinary) Byte(i int) byte {
	return byte(n.Uint8(i))
}
