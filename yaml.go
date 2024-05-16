package object

import (
	"reflect"

	"github.com/spf13/viper"
)

func NewWithYaml[T Struct](yaml *viper.Viper) *T {
	return SetYaml(NewWithDefault[T](), yaml)
}

func SetYaml[T Struct](obj *T, yaml *viper.Viper) *T {
	setYaml(reflect.ValueOf(obj), yaml)

	return obj
}

func setYaml(obj reflect.Value, yaml *viper.Viper) {
	// nil pointer
	if obj.Kind() == reflect.Ptr && obj.IsNil() {
		return
	}

	if obj.Type().Kind() == reflect.Ptr {
		setYaml(obj.Elem(), yaml)
		return
	}

	if obj.Type().Kind() != reflect.Struct {
		return
	}

	if yaml == nil || len(yaml.AllKeys()) == 0 {
		return
	}

	obj_type := obj.Type()
	for i, num := 0, obj_type.NumField(); i < num; i++ {
		field := obj_type.Field(i)
		// *struct{}
		if field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct {
			if f := obj.FieldByName(field.Name); !f.IsNil() { // init field value
				key, _ := field.Tag.Lookup("yaml")
				setYaml(obj.FieldByName(field.Name).Elem(), getSubViper(yaml, key))
			} else if f.CanSet() { // unint field value
				obj.FieldByName(field.Name).Set(reflect.New(field.Type.Elem()))

				key, _ := field.Tag.Lookup("yaml")
				setYaml(obj.FieldByName(field.Name).Elem(), getSubViper(yaml, key))
			}
		} else if field.Type.Kind() == reflect.Struct {
			key, _ := field.Tag.Lookup("yaml")
			setYaml(obj.FieldByName(field.Name), getSubViper(yaml, key))
		} else if key, ok := field.Tag.Lookup("yaml"); ok && yaml.IsSet(key) {
			if setValue, ok := setValueMap[field.Type.Kind()]; ok {
				setValue(obj.FieldByName(field.Name), yaml.GetString(key))
			}
		}
	}
}

func getSubViper(yaml *viper.Viper, key string) *viper.Viper {
	if yaml == nil || key == "" || !yaml.IsSet(key) {
		return yaml
	}

	return yaml.Sub(key)
}
