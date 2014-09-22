package reallyunsafe

import "unsafe"

var (
	dummy  = 0x01020304
	dummyB = (*(*[4]byte)(unsafe.Pointer(&dummy)))[0]
)

// These variables are read only
var (
	IsLittleEndian = dummyB == 0x04
	IsBigEndian    = dummyB == 0x01
	IsMiddleEndian = dummyB == 0x02
)
