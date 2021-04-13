package value

import (
	"reflect"
	"strconv"
	"strings"
)

func SetField(d interface{}, value string, path string) bool {
	v := reflect.ValueOf(d).Elem()
	for _, key := range strings.Split(path, `.`) {
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		switch v.Type().Kind() {
		case reflect.Struct:
			v = v.FieldByName(key)
		case reflect.Slice:
			if i, err := strconv.Atoi(key); err == nil {
				v = v.Index(i)
			} else {
				return false
			}
		case reflect.Map:
			v = v.MapIndex(reflect.ValueOf(key))
		}
	}

	if !v.CanSet() {
		return false
	}

	iValue, err := Convert(v.Type().Name(), value)
	if err != nil {
		return false
	}

	rvv := reflect.ValueOf(iValue)
	if v.Type() != rvv.Type() {
		return false
	}
	v.Set(rvv)

	return true
}

func GetField(d interface{}, path string) interface{} {
	v := reflect.ValueOf(d).Elem()
	for _, key := range strings.Split(path, `.`) {
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if v.Kind() == reflect.Invalid {
			return nil
		}

		switch v.Type().Kind() {
		case reflect.Struct:
			v = v.FieldByName(key)
		case reflect.Slice:
			if i, err := strconv.Atoi(key); err == nil {
				v = v.Index(i)
			} else {
				return false
			}
		case reflect.Map:
			v = v.MapIndex(reflect.ValueOf(key))
		}
	}

	if v.Kind() == reflect.Invalid {
		return nil
	}

	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Bool:
		return v.Bool()
	case reflect.Struct, reflect.Interface:
		return v.Interface()
	}

	return nil
}
