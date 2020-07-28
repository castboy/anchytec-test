package constant

type ErrCode uint
type ErrCodeSub uint
type ErrCodeSubSub uint

const (
	SystemErr ErrCode = iota + 1
	TradeErr
	MysqlErr
	KafkaErr
	ArgsErr
	AssertErr
	UnknowErr
)

var ErrCodeMsg = map[ErrCode]string{
	SystemErr: "System",
	TradeErr:  "Trade",
	MysqlErr:  "Mysql",
	KafkaErr:  "Kafka",
	ArgsErr:   "Args",
	AssertErr: "Assert",
	UnknowErr: "Unknow",
}

