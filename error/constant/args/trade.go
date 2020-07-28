package args

import "anchytec/error/constant"

type TradeArg constant.ErrCodeSubSub

const (
	TypeErr TradeArg = iota + 1
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

var TradeArgsMsg = map[TradeArg]string{
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

