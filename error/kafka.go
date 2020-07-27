package error

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
	base.init(errOrigin, mysqlErr, errCodeSub(opType), errCodeSubSub(opObj), comment...)

	return &KafkaErrer{base}
}

func (i *KafkaErrer) getErrCodeSubMsg() string {
	return kafkaOpTypeMsg[kafkaOpType(i.codeSub)]
}

func (i *KafkaErrer) getErrCodeSubSubMsg() string {
	return kafkaOpObjMsg[kafkaOpObj(i.codeSubSub)]
}

func (i *KafkaErrer) Error() string {
	return encodeErrMsg(i)
}