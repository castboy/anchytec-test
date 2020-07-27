package error

type tradeArg errCodeSubSub

const (
	TypeErr tradeArg = iota + 1
	LoginErr
	TicketErr
	CmdErr
	SymbolErr
	LeverageErr
	VolumeErr
	TpErr
	SlErr
	CommentErr
	ClientPriceErr
	ClientTimeErr
	PendingPriceErr
	ExpirationErr
	AmountErr
)

var tradeArgsMsg = map[tradeArg]string{
	TypeErr:         "type",
	LoginErr:        "login",
	TicketErr:       "ticket",
	CmdErr:          "cmd",
	SymbolErr:       "symbol",
	LeverageErr:     "leverage",
	VolumeErr:       "volume",
	TpErr:           "tp",
	SlErr:           "sl",
	CommentErr:      "comment",
	ClientPriceErr:  "client-price",
	ClientTimeErr:   "client-time",
	PendingPriceErr: "pending-price",
	ExpirationErr:   "expiration",
	AmountErr:       "amount",
}

type TradeArgsErrer struct {
	*baseErrer
}

func NewTradeArgsErrer(errOrigin error, reqType argsType, tradeArg tradeArg, comment ...Comment) *TradeArgsErrer {
	base := newErr()
	base.setCode(argsErr)
	base.setCodeSub(errCodeSub(reqType))
	base.setCodeSubSub(errCodeSubSub(tradeArg))
	base.setOriginErr(errOrigin)

	f := getCallFunc(newDeep)
	base.appendCallFunc(f)

	base.appendComment(comment...)

	return &TradeArgsErrer{base}
}

func (i *TradeArgsErrer) getErrCodeSubMsg() string {
	return argsTypeMsg[argsType(i.codeSub)]
}

func (i *TradeArgsErrer) getErrCodeSubSubMsg() string {
	return tradeArgsMsg[tradeArg(i.codeSubSub)]
}

func (i *TradeArgsErrer) Error() string {
	return encodeErrMsg(i)
}