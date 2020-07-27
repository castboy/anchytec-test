package error

import (
	"fmt"
	"testing"
)

func TestNewMysqlErrer(t *testing.T) {
	errOrigin := fmt.Errorf("Error 1062: Duplicate entry '1995623428' for key 'PRIMARY'")
	err := NewMysqlErrer(errOrigin, Insert, Order)
	fmt.Println(err.Error())
}
