package core

import (
	"errors"
	"fmt"
	"reflect"
)

func ShouldBeStruct(d reflect.Type) error {
	td := d.Elem()
	if td.Kind() != reflect.Struct {
		errStr := fmt.Sprintf("Input should be %v, found %v", reflect.Struct, td.Kind())
		return errors.New(errStr)
	}
	return nil
}
