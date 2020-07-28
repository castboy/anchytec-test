package error

import (
	"anchytec/error/constant"
	"anchytec/error/constant/kafka"
)

type KafkaErrer struct {
	*baseErrer
}

func NewKafkaErrer(errOrigin error, opType kafka.KafkaOpType, opObj kafka.KafkaOpObj, comment ...Comment) *KafkaErrer {
	base := newErr()
	base.init(errOrigin, constant.MysqlErr, constant.ErrCodeSub(opType), constant.ErrCodeSubSub(opObj), comment...)

	return &KafkaErrer{base}
}

func (i *KafkaErrer) getErrCodeSubMsg() string {
	return kafka.KafkaOpTypeMsg[kafka.KafkaOpType(i.codeSub)]
}

func (i *KafkaErrer) getErrCodeSubSubMsg() string {
	return kafka.KafkaOpObjMsg[kafka.KafkaOpObj(i.codeSubSub)]
}

func (i *KafkaErrer) Error() string {
	return encodeErrMsg(i)
}