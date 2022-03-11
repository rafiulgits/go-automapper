package automapper

import "reflect"

func Map(src interface{}, dst interface{}, args ...interface{}) {
	_mapLoose(src, dst)

	if len(args) > 0 {
		arg := reflect.TypeOf(args[0])
		if arg.Kind() != reflect.Func {
			// first arg is not a function
			return
		}
		if arg.NumIn() != 2 {
			// function must containe 2 params as arg
			return
		}

		if arg.In(0) != reflect.TypeOf(src) || arg.In(1) != reflect.TypeOf(dst) {
			// function params didn't matched with src and dst
			return
		}

		fun := reflect.ValueOf(args[0])
		fun.Call([]reflect.Value{reflect.ValueOf(src), reflect.ValueOf(dst)})
	}
}
