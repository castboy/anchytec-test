package test

import (
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
