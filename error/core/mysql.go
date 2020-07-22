package core

import (
	"fmt"
)

type mysqlOpType errCodeSub
type mysqlOpObj errCodeSubSub

const (
	Insert mysqlOpType = iota + 1
	Delete
	Update
	Search
	Other
)

var mysqlOpTypeMsg = map[mysqlOpType]string{
	Insert: "Insert",
	Delete: "Delete",
	Update: "Update",
	Search: "Search",
	Other:  "Other",
}

const (
	Order mysqlOpObj = iota + 1
	Symbol
	Source
)

var mysqlOpObjMsg = map[mysqlOpObj]string{
	Order:  "Order",
	Symbol: "Symbol",
	Source: "Source",
}

type MysqlErrer struct {
	*baseErrer
}

func NewMysqlErrer(opType mysqlOpType, opObj mysqlOpObj, comment ...Comment) *MysqlErrer {
	base := newErr(mysqlErr, errCodeSub(opType))
	base.setCodeSubSub(errCodeSubSub(opObj))
	base.appendComment(comment...)

	return &MysqlErrer{base}
}

func (i *MysqlErrer) encodeErrCode() string {
	return fmt.Sprintf("%s, %s, %s", errCodeMsg[i.code], mysqlOpTypeMsg[mysqlOpType(i.codeSub)], mysqlOpObjMsg[mysqlOpObj(i.codeSubSub)])
}

func (i *MysqlErrer) Error() string {
	return encodeError(i.encodeErrCode(), i.baseErrer.encodeComment(), i.originErr)
}
