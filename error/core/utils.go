package core

import (
	"runtime"
)

func getCallFunc() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	return runtime.FuncForPC(pc[0]).Name()
}
