package core

const (
	trade errCodeSub = iota + 1
)

var errCodeMsgArgs = map[errCodeSub]string{
	trade: "trade",
}

const (
	tradeTypeErr errCodeSubSub = iota + 1
	tradeLoginErr
	tradeTicketErr
	tradeCmdErr
	tradeSymbolErr
	tradeLeverageErr
	tradeVolumeErr
	tradeTpErr
	tradeSlErr
	tradeCommentErr
	tradeClientPriceErr
	tradeClientTimeErr
	tradePendingPriceErr
	tradeExpirationErr
	tradeAmountErr
)

var errCodeMsgArgsTrade = map[errCodeSubSub]string{
	tradeTypeErr:         "type",
	tradeLoginErr:        "login",
	tradeTicketErr:       "ticket",
	tradeCmdErr:          "cmd",
	tradeSymbolErr:       "symbol",
	tradeLeverageErr:     "leverage",
	tradeVolumeErr:       "volume",
	tradeTpErr:           "tp",
	tradeSlErr:           "sl",
	tradeCommentErr:      "comment",
	tradeClientPriceErr:  "client-price",
	tradeClientTimeErr:   "client-time",
	tradePendingPriceErr: "pending-price",
	tradeExpirationErr:   "expiration",
	tradeAmountErr:       "amount",
}

type argsErrer struct {
	*baseErrer
}

func
