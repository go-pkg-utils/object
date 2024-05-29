package object

import (
	"reflect"

	"github.com/go-pkg-utils/object/internal"
)

func NewWithDefaults[T internal.Struct]() *T {
	return SetDefaults(new(T))
}

func SetDefaults[T internal.Struct](obj *T) *T {
	setDefaults(reflect.ValueOf(obj))

	return obj
}

func setDefaults(obj reflect.Value) {
	// nil pointer
	if obj.Kind() == reflect.Ptr && obj.IsNil() {
		return
	}

	if obj.Type().Kind() == reflect.Ptr {
		setDefaults(obj.Elem())
		return
	}

	if obj.Type().Kind() != reflect.Struct {
		return
	}

	obj_type := obj.Type()
	for i, num := 0, obj_type.NumField(); i < num; i++ {
		field := obj_type.Field(i)
		// *struct{}
		if field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct {
			if f := obj.FieldByName(field.Name); !f.IsNil() { // init field value
				setDefaults(obj.FieldByName(field.Name).Elem())
			} else if f.CanSet() { // unint field value
				obj.FieldByName(field.Name).Set(reflect.New(field.Type.Elem()))
				setDefaults(obj.FieldByName(field.Name).Elem())
			}
		} else if field.Type.Kind() == reflect.Struct {
			setDefaults(obj.FieldByName(field.Name))
		} else if value, ok := field.Tag.Lookup("default"); ok {
			if setValue, ok := internal.SetValueMap[field.Type.Kind()]; ok {
				setValue(obj.FieldByName(field.Name), value)
			}
		}
	}
}
