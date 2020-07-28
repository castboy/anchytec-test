package system

import "anchytec/error/constant"

type Type constant.ErrCodeSub

const (
	SystemErr Type = iota + 1 // package any panic.
	SystemBusy // to reject user request.
	JobTimeout
)

var SystemTypeMsg = map[Type]string {
	SystemErr: "System Error",
	SystemBusy: "System Busy",
	JobTimeout: "Job Timeout",
}