// Copyright (c) 2015 Peter Strøiman, distributed under the MIT license

// Package automapper provides support for mapping between two different types
// with compatible fields. The intended application for this is when you use
// one set of types to represent DTOs (data transfer objects, e.g. json data),
// and a different set of types internally in the application. Using this
// package can help converting from one type to another.
//
// This package uses reflection to perform mapping which should be fine for
// all but the most demanding applications.
package automapper

import (
	"reflect"
	"time"
)

func _mapLoose(source, dest interface{}) {
	var dstType = reflect.TypeOf(dest)
	if dstType.Kind() != reflect.Ptr {
		panic("destination must be a pointer type")
	}
	var srcVal = reflect.ValueOf(source)
	var dstVal = reflect.ValueOf(dest).Elem()
	mapValues(srcVal, dstVal, true)
}

func mapValues(srcVal, dstVal reflect.Value, loose bool) {
	dstType := dstVal.Type()
	if dstType.Kind() == reflect.Struct {
		if srcVal.Type().Kind() == reflect.Ptr {
			if srcVal.IsNil() {
				// If source is nil, it maps to an empty struct
				srcVal = reflect.New(srcVal.Type().Elem())
			}
			srcVal = srcVal.Elem()
		}
		// If destination type if time object then try to map time object only
		if valueIsTimeObject(dstVal) {
			mapTime(srcVal, dstVal)
			return
		}
		for i := 0; i < dstVal.NumField(); i++ {
			mapField(srcVal, dstVal, i, loose)
		}
	} else if dstType == srcVal.Type() {
		dstVal.Set(srcVal)
	} else if dstType.Kind() == reflect.Ptr {
		if valueIsNil(srcVal) {
			return
		}
		val := reflect.New(dstType.Elem())
		mapValues(srcVal, val.Elem(), loose)
		dstVal.Set(val)
	} else if dstType.Kind() == reflect.Slice {
		mapSlice(srcVal, dstVal, loose)
	} else {
		//mapping currently not supported

	}
}

func mapSlice(srcVal, dstVal reflect.Value, loose bool) {
	dstType := dstVal.Type()
	length := srcVal.Len()
	target := reflect.MakeSlice(dstType, length, length)
	for j := 0; j < length; j++ {
		val := reflect.New(dstType.Elem()).Elem()
		mapValues(srcVal.Index(j), val, loose)
		target.Index(j).Set(val)
	}

	if length == 0 {
		verifyArrayTypesAreCompatible(srcVal, dstVal, loose)
	}
	dstVal.Set(target)
}

func verifyArrayTypesAreCompatible(srcVal, dstVal reflect.Value, loose bool) {
	dummyDest := reflect.New(reflect.PtrTo(dstVal.Type()))
	dummySource := reflect.MakeSlice(srcVal.Type(), 1, 1)
	mapValues(dummySource, dummyDest.Elem(), loose)
}

func mapField(source, dstVal reflect.Value, i int, loose bool) {
	dstType := dstVal.Type()
	fieldName := dstType.Field(i).Name

	destField := dstVal.Field(i)
	if dstType.Field(i).Anonymous {
		mapValues(source, destField, loose)
	} else {
		if valueIsContainedInNilEmbeddedType(source, fieldName) {
			return
		}
		sourceField := source.FieldByName(fieldName)
		if (sourceField == reflect.Value{}) {
			if loose {
				return
			}
			if destField.Kind() == reflect.Struct {
				mapValues(source, destField, loose)
				return
			} else {
				for i := 0; i < source.NumField(); i++ {
					if source.Field(i).Kind() != reflect.Struct {
						continue
					}
					if sourceField = source.Field(i).FieldByName(fieldName); (sourceField != reflect.Value{}) {
						break
					}
				}
			}
		}
		mapValues(sourceField, destField, loose)
	}
}

func valueIsNil(value reflect.Value) bool {
	return value.Type().Kind() == reflect.Ptr && value.IsNil()
}

func valueIsContainedInNilEmbeddedType(source reflect.Value, fieldName string) bool {
	structField, _ := source.Type().FieldByName(fieldName)
	ix := structField.Index
	if len(structField.Index) > 1 {
		parentField := source.FieldByIndex(ix[:len(ix)-1])
		if valueIsNil(parentField) {
			return true
		}
	}
	return false
}

func valueIsTimeObject(value reflect.Value) bool {
	_, ok := value.Interface().(time.Time)
	return ok
}

func mapTime(srcVal, dstVal reflect.Value) {
	switch srcVal.Type().Kind() {
	case reflect.String:
		parsedTime, err := time.Parse(time.RFC3339, srcVal.String())
		if err == nil {
			dstVal.Set(reflect.ValueOf(parsedTime))

		}

	case reflect.Struct:
		dstVal.Set(srcVal)
	}
}
