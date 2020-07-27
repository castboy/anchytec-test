package error

import (
	"runtime"
)

func getCallFunc() string {
	pc := make([]uintptr, 1)
	runtime.Callers(1, pc)
	return runtime.FuncForPC(pc[0]).Name()
}
