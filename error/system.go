package error

import (
	"anchytec/error/constant"
	"anchytec/error/constant/system"
)

type SystemErrer struct {
	*baseErrer
}

func NewSystemErrer(errOrigin error, Type system.Type) *SystemErrer {
	base := newErr()
	base.init(errOrigin, constant.SystemErr, constant.ErrCodeSub(Type), constant.ErrCodeSubSub(0))

	return &SystemErrer{base}
}

func (i *SystemErrer) getErrCodeSubMsg() string {
	return system.SystemTypeMsg[system.Type(i.codeSub)]
}

func (i *SystemErrer) getErrCodeSubSubMsg() string {
	return ""
}

func (i *SystemErrer) Error() string {
	return encodeErrMsg(i)
}

