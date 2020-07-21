package core

import (
	"fmt"
)

const (
	MysqlInsert errCodeSub = iota + 1
	mysqlDelete
	mysqlUpdate
	mysqlSearch
	mysqlOther
)

var errCodeMsgMysql = map[errCodeSub]string{
	MysqlInsert: "Insert",
	mysqlDelete: "Delete",
	mysqlUpdate: "Update",
	mysqlSearch: "Search",
	mysqlOther:  "Other",
}

const (
	MysqlInsertOrder errCodeSubSub = iota + 1
)

var errCodeMsgMysqlInsert = map[errCodeSubSub]string{
	MysqlInsertOrder: "Order",
}

const (
	mysqlDeleteOrder errCodeSubSub = iota + 1
)

var errCodeMsgMysqlDelete = map[errCodeSubSub]string{
	mysqlDeleteOrder: "Order",
}

const (
	mysqlUpdateOrder errCodeSubSub = iota + 1
)

var errCodeMsgMysqlUpdate = map[errCodeSubSub]string{
	mysqlUpdateOrder: "Order",
}

const (
	mysqlSearchOrder errCodeSubSub = iota + 1
)

var errCodeMsgMysqlSearch = map[errCodeSubSub]string{
	mysqlSearchOrder: "Order",
}

const (
	mysqlOther1 errCodeSubSub = iota + 1
	mysqlOther2
)

var errCodeMsgMysqlOther = map[errCodeSubSub]string{
	mysqlOther1: "Order1",
	mysqlOther2: "Order2",
}

type MysqlErrer struct {
	*baseErrer
}

func NewMysqlErrer(subCode errCodeSub, subSubCode errCodeSubSub) *MysqlErrer {
	base := newErr(mysqlErr, subCode)
	base.setCodeSubSub(subSubCode)

	return &MysqlErrer{base}
}

func (i *MysqlErrer) getErrCodeMsgMysqlSub() map[errCodeSubSub]string {
	switch i.codeSub {
	case MysqlInsert:
		return errCodeMsgMysqlInsert
	case mysqlDelete:
		return errCodeMsgMysqlDelete
	case mysqlUpdate:
		return errCodeMsgMysqlUpdate
	case mysqlSearch:
		return errCodeMsgMysqlSearch
	case mysqlOther:
		return errCodeMsgMysqlOther
	default:
		panic(fmt.Sprintf("invalid codeSub: %d", i.codeSub))
	}

	return nil
}

func (i *MysqlErrer) Error() string {
	errCodeMsg := fmt.Sprintf("%s, %s, %s", errCodeMsg[i.code], errCodeMsgMysql[i.codeSub], i.getErrCodeMsgMysqlSub()[i.codeSubSub])

	comment := ""
	for j := range i.comment {
		comment = fmt.Sprintf(" %s, %+v,", comment, i.comment[j])
	}

	return fmt.Sprintf("ERROR_CODE_MSG: %s, ERROR_ORIGIN: %v, COMMENT: %s", errCodeMsg, i.originErr, comment)
}
