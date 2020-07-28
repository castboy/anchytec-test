package args

import "anchytec/error/constant"

type ArgsType constant.ErrCodeSub

const (
	Trade ArgsType = iota + 1
)

var ArgsTypeMsg = map[ArgsType]string{
	Trade: "Trade",
}