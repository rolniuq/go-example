package day11

import (
	"errors"
	"reflect"
)

type Day11 struct {
}

func (d *Day11) CopyStruct(dst, src any) error {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)

	if dstVal.Kind() != reflect.Pointer || dstVal.Elem().Kind() != reflect.Struct {
		return errors.New("dst must be pointer to struct")
	}

	if srcVal.Kind() != reflect.Pointer || dstVal.Elem().Kind() != reflect.Struct {
		return errors.New("src must be pointer to struct")
	}

	for i := range dstVal.NumField() {
		dstField := dstVal.Field(i)
		dstFieldType := dstField.Type()

		if !dstField.CanSet() {
			continue
		}

		srcField := srcVal.FieldByName(dstFieldType.Name())
		if !srcField.IsValid() {
			continue
		}

		if srcField.Type() != dstFieldType {
			continue
		}

		dstField.Set(srcField)
	}

	return nil
}

func (d *Day11) Exec() {

}
