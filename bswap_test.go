package reallyunsafe

import "testing"

func SwapUint64Go(x uint64) uint64 {
	x = (0xff00ff00ff00ff & (x >> 8)) | ((0xff00ff00ff00ff & x) << 8)
	x = (0xffff0000ffff & (x >> 16)) | ((0xffff0000ffff & x) << 16)
	return (0xffffffff & (x >> 32)) | ((0xffffffff & x) << 32)
}

func BenchmarkSwapAsm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := SwapUint64(1985)
		if SwapUint64(i) != 1985 {
			panic("!1985")
		}
	}
}

func BenchmarkSwapGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := SwapUint64Go(1985)
		if SwapUint64Go(i) != 1985 {
			panic("!1985")
		}
	}
}
