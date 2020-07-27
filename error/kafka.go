package error

import "fmt"

type kafkaOpType errCodeSub
type kafkaOpObj errCodeSubSub

const (
	produce kafkaOpType = iota + 1
	consume
)

var kafkaOpTypeMsg = map[kafkaOpType]string{
	produce: "produce",
	consume: "consume",
}

const (
	order kafkaOpObj = iota + 1
	event
)

var kafkaOpObjMsg = map[kafkaOpObj]string{
	order: "order",
	event: "event",
}

type KafkaErrer struct {
	*baseErrer
}

func NewKafkaErrer(errOrigin error, opType kafkaOpType, opObj kafkaOpObj, comment ...Comment) *KafkaErrer {
	base := newErr()
	base.setCode(kafkaErr)
	base.setCodeSub(errCodeSub(opType))
	base.setCodeSubSub(errCodeSubSub(opObj))
	base.setOriginErr(errOrigin)

	f := getCallFunc(newDeep)
	base.appendCallFunc(f)

	base.appendComment(comment...)

	return &KafkaErrer{base}
}

func (i *KafkaErrer) encodeErrCode() string {
	return fmt.Sprintf("%s, %s, %s", errCodeMsg[i.code], kafkaOpTypeMsg[kafkaOpType(i.codeSub)], kafkaOpObjMsg[kafkaOpObj(i.codeSubSub)])
}

func (i *KafkaErrer) Error() string {
	return encodeErrMsg(i)
}