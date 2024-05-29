package internal

import (
	"reflect"
	"strconv"
	"unsafe"
)

type Struct = interface{}

var SetValueMap = map[reflect.Kind]func(field reflect.Value, value string){
	reflect.String: func(field reflect.Value, value string) {
		if v, err := value, error(nil); err == nil {
			*(*string)(unsafe.Pointer(field.UnsafeAddr())) = v
		}
	},
	reflect.Bool: func(field reflect.Value, value string) {
		if v, err := strconv.ParseBool(value); err == nil {
			*(*bool)(unsafe.Pointer(field.UnsafeAddr())) = v
		}
	},
	reflect.Int: func(field reflect.Value, value string) {
		if v, err := strconv.ParseInt(value, 10, 0); err == nil {
			*(*int)(unsafe.Pointer(field.UnsafeAddr())) = int(v)
		}
	},
	reflect.Int8: func(field reflect.Value, value string) {
		if v, err := strconv.ParseInt(value, 10, 8); err == nil {
			*(*int8)(unsafe.Pointer(field.UnsafeAddr())) = int8(v)
		}
	},
	reflect.Int16: func(field reflect.Value, value string) {
		if v, err := strconv.ParseInt(value, 10, 16); err == nil {
			*(*int16)(unsafe.Pointer(field.UnsafeAddr())) = int16(v)
		}
	},
	reflect.Int32: func(field reflect.Value, value string) {
		if v, err := strconv.ParseInt(value, 10, 32); err == nil {
			*(*int32)(unsafe.Pointer(field.UnsafeAddr())) = int32(v)
		}
	},
	reflect.Int64: func(field reflect.Value, value string) {
		if v, err := strconv.ParseInt(value, 10, 64); err == nil {
			*(*int64)(unsafe.Pointer(field.UnsafeAddr())) = v
		}
	},
	reflect.Uint: func(field reflect.Value, value string) {
		if v, err := strconv.ParseUint(value, 10, 0); err == nil {
			*(*uint)(unsafe.Pointer(field.UnsafeAddr())) = uint(v)
		}
	},
	reflect.Uint8: func(field reflect.Value, value string) {
		if v, err := strconv.ParseUint(value, 10, 8); err == nil {
			*(*uint8)(unsafe.Pointer(field.UnsafeAddr())) = uint8(v)
		}
	},
	reflect.Uint16: func(field reflect.Value, value string) {
		if v, err := strconv.ParseUint(value, 10, 16); err == nil {
			*((*uint16)(unsafe.Pointer(field.UnsafeAddr()))) = uint16(v)
		}
	},
	reflect.Uint32: func(field reflect.Value, value string) {
		if v, err := strconv.ParseUint(value, 10, 32); err == nil {
			*(*uint32)(unsafe.Pointer(field.UnsafeAddr())) = uint32(v)
		}
	},
	reflect.Uint64: func(field reflect.Value, value string) {
		if v, err := strconv.ParseUint(value, 10, 64); err == nil {
			*(*uint64)(unsafe.Pointer(field.UnsafeAddr())) = v
		}
	},
	reflect.Float32: func(field reflect.Value, value string) {
		if v, err := strconv.ParseFloat(value, 32); err == nil {
			*(*float32)(unsafe.Pointer(field.UnsafeAddr())) = float32(v)
		}
	},
	reflect.Float64: func(field reflect.Value, value string) {
		if v, err := strconv.ParseFloat(value, 64); err == nil {
			*(*float64)(unsafe.Pointer(field.UnsafeAddr())) = v
		}
	},
	reflect.Complex64: func(field reflect.Value, value string) {
		if v, err := strconv.ParseComplex(value, 64); err == nil {
			*(*complex64)(unsafe.Pointer(field.UnsafeAddr())) = complex64(v)
		}
	},
	reflect.Complex128: func(field reflect.Value, value string) {
		if v, err := strconv.ParseComplex(value, 128); err == nil {
			*(*complex128)(unsafe.Pointer(field.UnsafeAddr())) = v
		}
	},
}
