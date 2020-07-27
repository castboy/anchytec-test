package error

import (
	"encoding/json"
)

type baseErrer struct {
	code       errCode
	codeSub    errCodeSub
	codeSubSub errCodeSubSub
	originErr  error
	callFunc   []string
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
	res := &baseErrer{}
	return res
}

func (me *baseErrer) init(errOrigin error, code errCode, codeSub errCodeSub, codeSubSub errCodeSubSub, comment ...Comment) {
	me.setCode(code)
	me.setCodeSub(codeSub)
	me.setCodeSubSub(codeSubSub)
	me.setOriginErr(errOrigin)

	f := getCallFunc(newDeep)
	me.appendCallFunc(f)

	me.appendComment(comment...)
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

func (me *baseErrer) appendCallFunc(f ...string) {
	me.callFunc = append(me.callFunc, f...)
}

func (me *baseErrer) AppendCallFunc() {
	f := getCallFunc(appendDeep)
	me.appendCallFunc(f)
}

func (me *baseErrer) appendComment(cmt ...Comment) {
	me.comment = append(me.comment, cmt...)
}

func (me *baseErrer) encodeCallFunc() string {
	l := len(me.callFunc)
	s := ""
	for i := l-1; i >= 0; i-- {
		s += me.callFunc[i]
		if i != 0 {
			s += "->"
		}
	}

	return s
}

func (me *baseErrer) encodeComment() string {
	comment, _ := json.Marshal(me.comment)
	return string(comment)
}

func (me *baseErrer) getOriginErr() error {
	return me.originErr
}

func (i *baseErrer) getErrCodeMsg() string {
	return errCodeMsg[i.code]
}