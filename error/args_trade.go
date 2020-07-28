package error

import (
	"anchytec/error/constant"
	"anchytec/error/constant/args"
)

type TradeArgsErrer struct {
	*baseErrer
}

func NewTradeArgsErrer(errOrigin error, argsType args.ArgsType, tradeArg args.TradeArg, comment ...Comment) *TradeArgsErrer {
	base := newErr()
	base.init(errOrigin, constant.MysqlErr, constant.ErrCodeSub(argsType), constant.ErrCodeSubSub(tradeArg), comment...)

	return &TradeArgsErrer{base}
}

func (i *TradeArgsErrer) getErrCodeSubMsg() string {
	return args.ArgsTypeMsg[args.ArgsType(i.codeSub)]
}

func (i *TradeArgsErrer) getErrCodeSubSubMsg() string {
	return args.TradeArgsMsg[args.TradeArg(i.codeSubSub)]
}

func (i *TradeArgsErrer) Error() string {
	return encodeErrMsg(i)
}