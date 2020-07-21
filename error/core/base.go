package core

import (
	"fmt"
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
	kafkaErr
	argsErr
	assertErr
	unknowErr
)

var errCodeMsg = map[errCode]string{
	systemErr: "System",
	tradeErr:  "Trade",
	mysqlErr:  "Mysql",
	kafkaErr:  "Kafka",
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

func (me *baseErrer) Error(errCodeMsg string) string {
	comment := ""
	for j := range me.comment {
		comment = fmt.Sprintf(" %s, %+v,", comment, me.comment[j])
	}

	return fmt.Sprintf("ERROR_CODE_MSG: %s, ERROR_ORIGIN: %v, COMMENT: %s", errCodeMsg, me.originErr, comment)
}
