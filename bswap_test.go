package reallyunsafe

import "testing"

func SwapUint64Go(x uint64) uint64 {
	x = (0xff00ff00ff00ff & (x >> 8)) | ((0xff00ff00ff00ff & x) << 8)
	x = (0xffff0000ffff & (x >> 16)) | ((0xffff0000ffff & x) << 16)
	return (0xffffffff & (x >> 32)) | ((0xffffffff & x) << 32)
}

func SwapUint32Go(x uint32) uint32 {
	return ((x << 24) & 0xff000000) |
		((x << 8) & 0x00ff0000) |
		((x >> 8) & 0x0000ff00) |
		((x >> 24) & 0x000000ff)
}

func BenchmarkSwapAsm32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := SwapUint32(1985)
		if v = SwapUint32(v); v != 1985 {
			b.Log("v", v)
		}
	}
}

func BenchmarkSwapAsm64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := SwapUint64(1985)
		if v = SwapUint64(v); v != 1985 {
			panic("!1985")
		}
	}
}

func BenchmarkSwapGo32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := SwapUint32Go(1985)
		if v = SwapUint32Go(v); v != 1985 {
			panic("!1985")
		}
	}
}

func BenchmarkSwapGo64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := SwapUint64Go(1985)
		if v = SwapUint64Go(v); v != 1985 {
			panic("!1985")
		}
	}
}
