package unsafe

import (
	"reflect"
	"testing"
	"unsafe"
)

func string2bytesUnsafe(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2stringUnsafe(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func string2bytes(s string) []byte {
	return []byte(s)
}

func bytes2string(b []byte) string {
	return string(b)
}

func BenchmarkString2bytesUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		string2bytesUnsafe("1234567")
	}
}

func BenchmarkBytes2stringUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes2stringUnsafe([]byte{49, 50, 51, 52, 53, 54, 55})
	}
}

func BenchmarkString2bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		string2bytes("1234567")
	}
}

func BenchmarkBytes2string(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes2string([]byte{49, 50, 51, 52, 53, 54, 55})
	}
}
