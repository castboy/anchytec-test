package mysql

import (
	"anchytec/error/constant"
)

type MysqlOpType constant.ErrCodeSub
type MysqlOpObj constant.ErrCodeSubSub

const (
	Insert MysqlOpType = iota + 1
	Delete
	Update
	Search
	Other
)

var MysqlOpTypeMsg = map[MysqlOpType]string{
	Insert: "Insert",
	Delete: "Delete",
	Update: "Update",
	Search: "Search",
	Other:  "Other",
}

const (
	Order MysqlOpObj = iota + 1
	Symbol
	Source
)

var MysqlOpObjMsg = map[MysqlOpObj]string{
	Order:  "Order",
	Symbol: "Symbol",
	Source: "Source",
}

