package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/shopspring/decimal"
	"log"
)

// decimal.Decimal type map table item type 'text' default. we can define certain type by set `xorm:"salary decimal(28,0) NOT NULL"`
// by NoAutoCondition(true), we can get a clean query sql, like:
// `SELECT `id`, `name`, `salary` FROM `person` WHERE id=1 AND name='wmg' LIMIT 1`, not
// `SELECT `id`, `name`, `salary` FROM `account4` WHERE id=1 AND name='wmg' AND `salary`='100' LIMIT 1`

type Person struct {
	Id     int64           `xorm:"id"`
	Name   string          `xorm:"name"`
	Salary decimal.Decimal `xorm:"salary decimal(28,0) NOT NULL"`
}

var x *xorm.Engine

func main() {
	x, err := xorm.NewEngine("mysql", "root:wang1234@/symbol?charset=utf8")

	if err = x.Sync2(new(Person)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}

	person := Person{
		Id:     1,
		Name:   "wmg",
		Salary: decimal.NewFromFloat(100),
	}

	i, err := x.Insert(&person)
	fmt.Println(i, err)

	a := &Person{}
	has, err := x.Where(map[string]interface{}{"id": 1, "name": "wmg"}).NoAutoCondition(true).Get(a)
	fmt.Println(has, a, err)

	has, err = x.Where("id=?", 1).NoAutoCondition(true).Get(a)
	fmt.Println(has, a, err)
}
