package error

type argsType errCodeSub

const (
	Trade argsType = iota + 1
)

var argsTypeMsg = map[argsType]string{
	Trade: "Trade",
}