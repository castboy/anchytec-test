package error

import (
	"anchytec/error/constant"
	"anchytec/error/constant/mysql"
)

type MysqlErrer struct {
	*baseErrer
}

func NewMysqlErrer(errOrigin error, opType mysql.MysqlOpType, opObj mysql.MysqlOpObj, comment ...Comment) *MysqlErrer {
	base := newErr()
	base.init(errOrigin, constant.MysqlErr, constant.ErrCodeSub(opType), constant.ErrCodeSubSub(opObj), comment...)

	return &MysqlErrer{base}
}

func (i *MysqlErrer) getErrCodeSubMsg() string {
	return mysql.MysqlOpTypeMsg[mysql.MysqlOpType(i.codeSub)]
}

func (i *MysqlErrer) getErrCodeSubSubMsg() string {
	return mysql.MysqlOpObjMsg[mysql.MysqlOpObj(i.codeSubSub)]
}

func (i *MysqlErrer) Error() string {
	return encodeErrMsg(i)
}
