package sizeof

import (
	"reflect"
)

func SizeOf(i interface{}) int64 {
	return sizeOf(reflect.ValueOf(i))
}

func sizeOf(valOf reflect.Value) int64 {
	var typeOf = valOf.Type()
	var typeKind = typeOf.Kind()
	switch typeKind {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64,
		reflect.Complex128, reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return int64(typeOf.Size())
	case reflect.Ptr:
		return sizeOfPtr(valOf)
	case reflect.Array:
		return int64(typeOf.Size())
	case reflect.Map:
		return sizeOfMap(valOf)
	case reflect.Slice:
		return sizeOfSlice(valOf)
	case reflect.String:
		return sizeOfString(valOf)
	case reflect.Struct:
		return sizeOfStruct(valOf)
	case reflect.Interface:
		return sizeOf(valOf.Elem()) + int64(valOf.Type().Size())
	}

	return 0
}

func sizeOfPtr(valOf reflect.Value) int64 {
	if valOf.IsNil() {
		return 0
	}

	return int64(valOf.Type().Size()) + sizeOf(valOf.Elem())
}

func sizeOfSlice(valOf reflect.Value) int64 {
	if valOf.IsNil() {
		return 0
	}

	var sizeOfArr int64
	for i := 0; i < valOf.Len(); i++ {
		sizeOfArr += sizeOf(valOf.Index(i))
	}

	return sizeOfArr + int64(valOf.Type().Size())
}

func sizeOfMap(valOf reflect.Value) int64 {
	if valOf.IsNil() {
		return 0
	}

	var sizeOfMap int64
	var iter = valOf.MapRange()
	for iter.Next() {
		sizeOfMap += sizeOf(iter.Key())
		sizeOfMap += sizeOf(iter.Value())
	}
	return sizeOfMap + int64(valOf.Type().Size())
}

func sizeOfString(valOf reflect.Value) int64 {
	return int64(valOf.Type().Size()) + int64(len(valOf.String())*8)
}

func sizeOfStruct(valOf reflect.Value) int64 {
	var sizeOfStruct int64
	var sizeOfFields int64
	for i := 0; i < valOf.NumField(); i++ {
		var field = valOf.Field(i)
		var size = sizeOf(field)
		sizeOfStruct += size
		if size == 0 && field.Kind() == reflect.Ptr && field.IsNil() {
			sizeOfStruct += int64(field.Type().Size())
		}
		sizeOfFields += int64(field.Type().Size())
	}
	var align = int64(valOf.Type().Size()) - sizeOfFields
	return sizeOfStruct + align
}
