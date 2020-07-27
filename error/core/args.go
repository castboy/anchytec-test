package core

type reqType errCodeSub
type tradeReq errCodeSubSub

const (
	Trade reqType = iota + 1
)

var reqTypeMsg = map[reqType]string{
	Trade: "Trade",
}

const (
	TypeErr tradeReq = iota + 1
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

var tradeArgsMsg = map[tradeReq]string{
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

func NewArgsErrer(opType reqType, subSub errCodeSubSub) *ArgsErrer {
	return &ArgsErrer{}
}
