package object

import (
	"bytes"
	"reflect"

	"github.com/spf13/viper"
)

func NewWithJson[T Struct](json string) *T {
	return SetJsonValue(NewWithDefault[T](), json)
}

func SetJsonValue[T Struct](obj *T, json string) *T {
	viper := viper.New()
	viper.SetConfigType("json")
	if err := viper.MergeConfig(bytes.NewReader([]byte(json))); err == nil {
		setJsonValue(reflect.ValueOf(obj), viper)
	}

	return obj
}

func setJsonValue(obj reflect.Value, viper *viper.Viper) {
	// nil pointer
	if obj.Kind() == reflect.Ptr && obj.IsNil() {
		return
	}

	if obj.Type().Kind() == reflect.Ptr {
		setJsonValue(obj.Elem(), viper)
		return
	}

	if obj.Type().Kind() != reflect.Struct {
		return
	}

	if viper == nil || len(viper.AllKeys()) == 0 {
		return
	}

	obj_type := obj.Type()
	for i, num := 0, obj_type.NumField(); i < num; i++ {
		field := obj_type.Field(i)
		// *struct{}
		if field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct {
			if f := obj.FieldByName(field.Name); !f.IsNil() { // init field value
				key, _ := field.Tag.Lookup("json")
				setJsonValue(obj.FieldByName(field.Name).Elem(), getSubViper(viper, key))
			} else if f.CanSet() { // unint field value
				obj.FieldByName(field.Name).Set(reflect.New(field.Type.Elem()))

				key, _ := field.Tag.Lookup("json")
				setJsonValue(obj.FieldByName(field.Name).Elem(), getSubViper(viper, key))
			}
		} else if field.Type.Kind() == reflect.Struct {
			key, _ := field.Tag.Lookup("json")
			setJsonValue(obj.FieldByName(field.Name), getSubViper(viper, key))
		} else if key, ok := field.Tag.Lookup("json"); ok && viper.IsSet(key) {
			if setValue, ok := setValueMap[field.Type.Kind()]; ok {
				setValue(obj.FieldByName(field.Name), viper.GetString(key))
			}
		}
	}
}
