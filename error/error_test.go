package error

import (
	"anchytec/error/constant/mysql"
	"fmt"
	"testing"
)

func TestNewMysqlErrer(t *testing.T) {
	err := NewMysqlErrerUpTwo()
	AppendCallFunc(err.(encodeError))

	fmt.Println(err.Error())
}

func NewMysqlErrerUpTwo() error {
	err := NewMysqlErrerUpOne()
	AppendCallFunc(err.(encodeError))

	fmt.Println(err.Error())

	return err
}

func NewMysqlErrerUpOne() error {
	errOrigin := fmt.Errorf("Error 1062: Duplicate entry '1995623428' for key 'PRIMARY'")
	err := NewMysqlErrer(errOrigin, mysql.Insert, mysql.Order)
	fmt.Println(err.Error())
	return err
}
