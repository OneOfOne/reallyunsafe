package reallyunsafe

import (
	"encoding/binary"
	"math/rand"
	"testing"
)

func Uint64A(p uintptr, i int) uint64

var testData []byte = make([]byte, 1024*1024*8)

func init() {
	for i := range testData {
		testData[i] = byte(rand.Intn(255))
	}
}

func BenchmarkReallyUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		br := NewNativeBinary(&testData)
		for y := 0; y < len(testData); y += 8 {
			_ = br.Uint64(y)
		}
	}
}

func BenchmarkBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for y := 0; y < len(testData); y += 8 {
			_ = binary.LittleEndian.Uint64(testData[y : y+8])
		}
	}
}
