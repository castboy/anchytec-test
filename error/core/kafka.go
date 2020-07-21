package core

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

func NewKafkaErrer(opType kafkaOpType, opObj kafkaOpObj) *KafkaErrer {
	base := newErr(kafkaErr, errCodeSub(opType))
	base.setCodeSubSub(errCodeSubSub(opObj))

	return &KafkaErrer{base}
}

func (i *KafkaErrer) Error() string {
	errCodeMsg := fmt.Sprintf("%s, %s, %s", errCodeMsg[i.code], kafkaOpTypeMsg[kafkaOpType(i.codeSub)], kafkaOpObjMsg[kafkaOpObj(i.codeSubSub)])

	return i.baseErrer.Error(errCodeMsg)
}