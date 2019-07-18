package test

import (
	"bou.ke/monkey"
	"github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMonkey(t *testing.T) {
	Convey("gomonkey", t, func() {
		patchGuard := monkey.Patch(target, func(a string) string {
			return "gomonkey"
		})

		s := target("a")
		So(s, ShouldEqual, "gomonkey")

		patchGuard.Unpatch()
		s = target("a")
		So(s, ShouldEqual, "a")
	})

	Convey("monkey", t, func() {
		patches := gomonkey.ApplyFunc(target, func(a string) string {
			return "monkey"
		})

		s := target("a")
		So(s, ShouldEqual, "monkey")

		patches.Reset()
		s = target("a")
		So(s, ShouldEqual, "a")
	})
}

func target(a string) string {
	return a
}

// note: go test gomonkey_monkey_test.go -gcflags=all=-l