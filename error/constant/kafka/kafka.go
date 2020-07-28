package kafka

import "anchytec/error/constant"

type KafkaOpType constant.ErrCodeSub
type KafkaOpObj constant.ErrCodeSubSub

const (
	Produce KafkaOpType = iota + 1
	Consume
)

var KafkaOpTypeMsg = map[KafkaOpType]string{
	Produce: "Produce",
	Consume: "Consume",
}

const (
	Order KafkaOpObj = iota + 1
	Event
)

var KafkaOpObjMsg = map[KafkaOpObj]string{
	Order: "Order",
	Event: "Event",
}

