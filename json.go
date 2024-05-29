package object

import (
	"bytes"
	"reflect"

	"github.com/go-pkg-utils/object/internal"
	"github.com/spf13/viper"
)

func NewWithJson[T internal.Struct](json string) *T {
	return SetJson(NewWithDefaults[T](), json)
}

func SetJson[T internal.Struct](obj *T, json string) *T {
	_json := viper.New()
	_json.SetConfigType("json")
	if err := _json.MergeConfig(bytes.NewReader([]byte(json))); err == nil {
		setJson(reflect.ValueOf(obj), _json)
	}

	return obj
}

func setJson(obj reflect.Value, json *viper.Viper) {
	// nil pointer
	if obj.Kind() == reflect.Ptr && obj.IsNil() {
		return
	}

	if obj.Type().Kind() == reflect.Ptr {
		setJson(obj.Elem(), json)
		return
	}

	if obj.Type().Kind() != reflect.Struct {
		return
	}

	if json == nil || len(json.AllKeys()) == 0 {
		return
	}

	obj_type := obj.Type()
	for i, num := 0, obj_type.NumField(); i < num; i++ {
		field := obj_type.Field(i)
		// *struct{}
		if field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct {
			if f := obj.FieldByName(field.Name); !f.IsNil() { // init field value
				key, _ := field.Tag.Lookup("json")
				setJson(obj.FieldByName(field.Name).Elem(), getSubViper(json, key))
			} else if f.CanSet() { // unint field value
				obj.FieldByName(field.Name).Set(reflect.New(field.Type.Elem()))

				key, _ := field.Tag.Lookup("json")
				setJson(obj.FieldByName(field.Name).Elem(), getSubViper(json, key))
			}
		} else if field.Type.Kind() == reflect.Struct {
			key, _ := field.Tag.Lookup("json")
			setJson(obj.FieldByName(field.Name), getSubViper(json, key))
		} else if key, ok := field.Tag.Lookup("json"); ok && json.IsSet(key) {
			if setValue, ok := internal.SetValueMap[field.Type.Kind()]; ok {
				setValue(obj.FieldByName(field.Name), json.GetString(key))
			}
		}
	}
}
