package automapper

import (
	"log"
	"reflect"
)

func Map(src interface{}, dst interface{}, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("automapper error:", r)
		}
	}()
	_mapLoose(src, dst)

	if len(args) == 0 {
		return
	}

	arg := reflect.TypeOf(args[0])
	if arg.Kind() != reflect.Func {
		// first arg is not a function
		return
	}

	// src:obj, dst:obj => (src:obj, dst:obj)
	// src:arr, dst:arr => (src:obj, dst:obj)

	srcType := reflect.TypeOf(src)
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// Slice  -> Slice mapping
	if srcType.Kind() == reflect.Slice {
		if arg.NumIn() != 3 {
			// function must contain 3 params as args (idx, src, dst)
			return
		}
		if arg.In(0).Kind() != reflect.Int {
			return
		}
		fun := reflect.ValueOf(args[0])
		for i := 0; i < srcVal.Len(); i++ {
			fun.Call([]reflect.Value{reflect.ValueOf(i), srcVal.Index(i), dstVal.Elem().Index(i)})
		}

		return
	}

	// Object to Object mapping
	if arg.NumIn() != 2 {
		// function must contain 2 params as args (src, dst)
		return
	}

	fun := reflect.ValueOf(args[0])
	fun.Call([]reflect.Value{srcVal, dstVal})
}
