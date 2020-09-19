package value

import (
	"reflect"
	"strconv"
	"strings"
)

func SetField(d interface{}, value string, path string) bool {
	v := reflect.ValueOf(d).Elem()
	for _, s := range strings.Split(path, `.`) {
		if i, err := strconv.Atoi(s); err == nil {
			v = v.Index(i)
		} else {
			v = v.FieldByName(s)
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
