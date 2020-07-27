package error

import (
	"runtime"
)

const (
	newDeep    = 4 //NewMysqlErrer()/NewKafkaErrer()...
	appendDeep = 4 // AppendCallFunc()
)

func getCallFunc(deep int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(deep, pc)
	return runtime.FuncForPC(pc[0]).Name()
}
