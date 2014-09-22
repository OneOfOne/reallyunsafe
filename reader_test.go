package reallyunsafe

import (
	"encoding/binary"
	"math/rand"
	"testing"
)

var testData []byte = make([]byte, 1024*1024*128)

func init() {
	for i := range testData {
		testData[i] = byte(rand.Intn(256) % 255)
	}
}
func BenchmarkBounds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		br := NewNativeBinary(&testData)
		for y := 0; ; y += 8 {
			if br.Uint64C(y) == 0 {
				break
			}
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
