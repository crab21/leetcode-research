package base

import (
	"testing"
	"unsafe"
)

var count = 10000

func BenchmarkByteTostring(b *testing.B) {
	b.Run("normal", func(b *testing.B) {
		ss := []byte("sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd")

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for s := 0; s < count; s++ {
				_ = string(ss)
			}
		}
	})

	b.Run("speed", func(b *testing.B) {
		ss := []byte("sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd")

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for s := 0; s < count; s++ {
				_ = *(*string)(unsafe.Pointer(&ss))
			}
		}
	})
}


func BenchmarkStringToByte(b *testing.B) {
	b.Run("normal", func(b *testing.B) {
		ss := "sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd"

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for s := 0; s < count; s++ {
				_ = []byte(ss)
			}
		}
	})

	b.Run("speed", func(b *testing.B) {
		ss := "sfdsfjsdkfjlsdfjsdlnslcnlsdfjlncnsdlkjflwefjwelfnl1831u482fjkencsd"

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for s := 0; s < count; s++ {
				_ = *(*[]byte)(unsafe.Pointer(&ss))
			}
		}
	})
}

/**
result:

goos: darwin
goarch: amd64
pkg: base
BenchmarkByteTostring/normal-8         	    3378	    337446 ns/op
BenchmarkByteTostring/speed-8          	  271171	      4278 ns/op
BenchmarkStringToByte/normal-8         	    2898	    394058 ns/op
BenchmarkStringToByte/speed-8          	  273120	      4271 ns/op
PASS

 */