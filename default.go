package object

import (
	"reflect"
)

func NewWithDefault[T Struct]() *T {
	return SetDefaultValue(new(T))
}

func SetDefaultValue[T Struct](obj *T) *T {
	setDefaultValue(reflect.ValueOf(obj))

	return obj
}

func setDefaultValue(obj reflect.Value) {
	// nil pointer
	if obj.Kind() == reflect.Ptr && obj.IsNil() {
		return
	}

	if obj.Type().Kind() == reflect.Ptr {
		setDefaultValue(obj.Elem())
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
				setDefaultValue(obj.FieldByName(field.Name).Elem())
			} else if f.CanSet() { // unint field value
				obj.FieldByName(field.Name).Set(reflect.New(field.Type.Elem()))
				setDefaultValue(obj.FieldByName(field.Name).Elem())
			}
		} else if field.Type.Kind() == reflect.Struct {
			setDefaultValue(obj.FieldByName(field.Name))
		} else if value, ok := field.Tag.Lookup("default"); ok {
			if setValue, ok := setValueMap[field.Type.Kind()]; ok {
				setValue(obj.FieldByName(field.Name), value)
			}
		}
	}
}
