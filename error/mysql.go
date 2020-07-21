package error

import (
	"trading-system/db"
	"anchytec/error/core"
)

func MysqlInsertOrder(err error, res []*db.TradeTransResult) *core.MysqlErrer {
	me := core.NewMysqlErrer(core.MysqlInsert, core.MysqlInsertOrder)
	me.SetOriginError(err)

	for i := range res {
		me.AppendComment(*res[i])
	}

	return me
}
