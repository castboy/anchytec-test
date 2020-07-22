package core

type ReqType errCodeSub
type TradeReq errCodeSubSub

const (
	Trade ReqType = iota + 1
)

var argsOpTypeMsg = map[ReqType]string{
	Trade: "Trade",
}

const (
	TypeErr TradeReq = iota + 1
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

var errCodeMsgArgsTrade = map[TradeReq]string{
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

type ArgsErrer struct {
	*baseErrer
}

func NewArgsErrer(opType ReqType, subSub errCodeSubSub) *ArgsErrer {
	return &ArgsErrer{}
}
