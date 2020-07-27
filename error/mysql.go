package error

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

func NewMysqlErrer(errOrigin error, opType mysqlOpType, opObj mysqlOpObj, comment ...Comment) *MysqlErrer {
	base := newErr()
	base.setCode(mysqlErr)
	base.setCodeSub(errCodeSub(opType))
	base.setCodeSubSub(errCodeSubSub(opObj))
	base.setOriginErr(errOrigin)

	f := getCallFunc(newDeep)
	base.appendCallFunc(f)

	base.appendComment(comment...)

	return &MysqlErrer{base}
}

func (i *MysqlErrer) getErrCodeSubMsg() string {
	return mysqlOpTypeMsg[mysqlOpType(i.codeSub)]
}

func (i *MysqlErrer) getErrCodeSubSubMsg() string {
	return mysqlOpObjMsg[mysqlOpObj(i.codeSubSub)]
}

func (i *MysqlErrer) Error() string {
	return encodeErrMsg(i)
}
