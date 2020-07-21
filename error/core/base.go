package core

import (
	"runtime"
)

type baseErrer struct {
	code       errCode
	codeSub    errCodeSub
	codeSubSub errCodeSubSub
	funcName   string
	originErr  error
	comment    []interface{}
}

type errCode uint
type errCodeSub uint
type errCodeSubSub uint

const (
	systemErr errCode = iota + 1
	tradeErr
	mysqlErr
	argsErr
	assertErr
	unknowErr
)

var errCodeMsg = map[errCode]string{
	systemErr: "System",
	tradeErr:  "Trade",
	mysqlErr:  "Mysql",
	argsErr:   "Args",
	assertErr: "Assert",
	unknowErr: "Unknow",
}

func newErr(code errCode, subCode errCodeSub) *baseErrer {
	err := &baseErrer{
		code:    code,
		codeSub: subCode,
	}

	err.setCallFuncName()

	return err
}

func (me *baseErrer) setCallFuncName() {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	me.funcName = runtime.FuncForPC(pc[0]).Name()
}

func (me *baseErrer) SetOriginError(err error) {
	me.originErr = err
}

func (me *baseErrer) setCodeSubSub(subSub errCodeSubSub) {
	me.codeSubSub = subSub
}

func (me *baseErrer) AppendComment(v interface{}) {
	me.comment = append(me.comment, v)
}
