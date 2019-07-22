package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/shopspring/decimal"
	"time"
)

// decimal.Decimal type map table item type 'text' default. we can define certain type by set `xorm:"salary decimal(28,0) NOT NULL"`
// by NoAutoCondition(true), we can get a clean query sql, like:
// `SELECT `id`, `name`, `salary` FROM `person` WHERE id=1 AND name='wmg' LIMIT 1`, not
// `SELECT `id`, `name`, `salary` FROM `account4` WHERE id=1 AND name='wmg' AND `salary`='100' LIMIT 1`

type Person struct {
	Id     int64           `xorm:"id autoincr"` // 这里不设置xorm tag,插入成功后，last id将赋给该字段
	Name   string          `xorm:"name"`
	Salary decimal.Decimal `xorm:"salary decimal(28,0) NOT NULL"`
	// `xorm:"xorm-type"`
	TimeStamp time.Time `xorm:"timestamp"`
	Time      time.Time `xorm:"time"`
	Date      time.Time `xorm:"date"`
	DateTime  time.Time `xorm:"datetime"`
	Created   time.Time `xorm:"created"`
	Updated   time.Time `xorm:"updated"`
}

var x *xorm.Engine

func main() {
	x, err := xorm.NewEngine("mysql", "root:wang1234@/symbol_test?charset=utf8")
	x.DatabaseTZ = time.UTC // local time saved as corresponding db utc time.
	x.TZLocation = time.Local // db utc time analyzed to corresponding local time.

	//if err = x.Sync2(new(Person)); err != nil {
	//	log.Fatalf("Fail to sync database: %v\n", err)
	//}

	person := Person{
		Name:      "wmg",
		Salary:    decimal.NewFromFloat(100),
		TimeStamp: time.Now().UTC(),
		Time:      time.Now().UTC(),
		Date:      time.Now().UTC(),
		DateTime:  time.Now().UTC(),
	}

	i, err := x.Insert(&person)
	fmt.Println(i, err, person.Id)

	// test Updated   time.Time `xorm:"updated"`
	time.Sleep(time.Second)
	person2 := Person{
		Name: "wmq",
	}
	_, err = x.Where(map[string]interface{}{"id": 1, "name": "wmg"}).Cols("name").Update(&person2)
	fmt.Println(err)

	a := &Person{}
	has, err := x.Where(map[string]interface{}{"id": i}).NoAutoCondition(true).Get(a)
	fmt.Println(has, err)
	fmt.Printf("%+v", a)
}
