package day11

import (
	"errors"
	"fmt"
	"log"
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

	if srcVal.Kind() != reflect.Pointer || srcVal.Elem().Kind() != reflect.Struct {
		return errors.New("src must be pointer to struct")
	}

	dstElem := dstVal.Elem()
	srcElem := srcVal.Elem()

	for i := 0; i < dstElem.NumField(); i++ {
		dstField := dstElem.Field(i)
		dstFieldInfo := dstElem.Type().Field(i)

		if !dstField.CanSet() || !dstFieldInfo.IsExported() {
			continue
		}

		srcField := srcElem.FieldByName(dstFieldInfo.Name)
		if !srcField.IsValid() || srcField.Type() != dstField.Type() {
			continue
		}

		dstField.Set(srcField)
	}

	return nil
}

func (d *Day11) Exec() {
	type A struct {
		X int
		y string
	}
	type B struct{ X int }

	a := &A{X: 42, y: "secret"}
	b := &B{}
	if err := d.CopyStruct(b, a); err != nil {
		log.Fatal(err)
	}

	fmt.Println(b.X)
}
