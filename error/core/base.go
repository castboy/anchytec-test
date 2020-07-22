package core

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type baseErrer struct {
	code       errCode
	codeSub    errCodeSub
	codeSubSub errCodeSubSub
	funcName   string
	originErr  error
	comment    []Comment
}

type Comment struct {
	k string
	v interface{}
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

func (me *baseErrer) setOriginError(err error) {
	me.originErr = err
}

func (me *baseErrer) setCodeSubSub(subSub errCodeSubSub) {
	me.codeSubSub = subSub
}

func (me *baseErrer) appendComment(cmt ...Comment) {
	me.comment = append(me.comment, cmt...)
}

func (me *baseErrer) encodeComment() string {
	comment, _ := json.Marshal(me.comment)
	return string(comment)
}

func encodeError(errCodeMsg, commentMsg string, errOrigin error, ) string {
	return fmt.Sprintf("ERROR_CODE_MSG: %s, ERROR_ORIGIN: %v, COMMENT: %s", errCodeMsg, errOrigin, commentMsg)
}
