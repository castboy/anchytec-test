package test

import (
	"fmt"
	"testing"
)

type Mark struct {
	start, end int
}

func beMarks() []Mark {
	length := 100000
	mks := make([]Mark, length)

	for i := 0; i < length; i++ {
		mks[i] = Mark{0, 1}
	}

	return mks
}

// lower
func Fa(mks []Mark) {
	count := 0

	for _, mk := range mks {
		count += mk.end - mk.start
	}
}

// faster
func Fb(mks []Mark) {
	count := 0

	for i := range mks {
		count += mks[i].end - mks[i].start
	}
}

func BenchmarkFa(b *testing.B) {
	mks := beMarks()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Fa(mks)
	}
}

func BenchmarkFb(b *testing.B) {
	mks := beMarks()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Fb(mks)
	}
}

var s = "éक्षिaπ囧éक्षिaπ囧éक्षिaπ囧éक्षिaπ囧éक्षिaπ囧éक्षिaπ囧éक्षिaπ囧éक्षिaπ囧éक्षिaπ囧"

func RangeByteLower() {
	for i := 0; i < len(s); i++ {
		fmt.Sprintf("The byte at index %v: 0x%x \n", i, s[i])
	}
}

func RangeByteFaster() {
	for i, b := range []byte(s) {
		fmt.Sprintf("The byte at index %v: 0x%x \n", i, b)
	}
}

func BenchmarkRangeByteLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeByteLower()
	}
}

func BenchmarkRangeByteFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeByteFaster()
	}
}
