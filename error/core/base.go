package core

import (
	"encoding/json"
)

type baseErrer struct {
	code       errCode
	codeSub    errCodeSub
	codeSubSub errCodeSubSub
	callFunc   string
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

func newErr() *baseErrer {
	return &baseErrer{
		callFunc: getCallFunc(),
	}
}

func (me *baseErrer) setCode(code errCode) {
	me.code = code
}

func (me *baseErrer) setCodeSub(sub errCodeSub) {
	me.codeSub = sub
}

func (me *baseErrer) setCodeSubSub(subSub errCodeSubSub) {
	me.codeSubSub = subSub
}

func (me *baseErrer) setOriginErr(err error) {
	me.originErr = err
}

func (me *baseErrer) appendComment(cmt ...Comment) {
	me.comment = append(me.comment, cmt...)
}

func (me *baseErrer) encodeComment() string {
	comment, _ := json.Marshal(me.comment)
	return string(comment)
}

func (me *baseErrer) getCallFunc() string {
	return me.callFunc
}

func (me *baseErrer) getOriginErr() error {
	return me.originErr
}
