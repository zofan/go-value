package value

import (
	"strconv"
	"strings"
)

func ToInt(v string) int64 {
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

func ToFloat(v string) float64 {
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func ToFloatSep(v, decimalSep, thousandsSep string) float64 {
	v = strings.TrimSpace(v)
	v = strings.ReplaceAll(v, thousandsSep, ``)
	v = strings.ReplaceAll(v, decimalSep, `.`)
	return ToFloat(v)
}

func ToBool(v string) bool {
	switch v {
	case `1`:
		return true
	case `y`, `Y`:
		return true
	case `true`, `True`, `TRUE`:
		return true
	case `yes`, `Yes`, `YES`:
		return true
	case `accept`, `Accept`, `ACCEPT`:
		return true
	case `ok`, `Ok`, `OK`:
		return true
	default:
		return false
	}
}

func Convert(t string, v string) (i interface{}, e error) {
	switch t {
	case `int`, `int8`, `int16`, `int32`, `int64`, `rune`:
		i, e = strconv.ParseInt(v, 10, 64)
	case `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `byte`:
		i, e = strconv.ParseUint(v, 10, 64)
	case `float32`, `float64`:
		i, e = strconv.ParseFloat(v, 64)
	case `bool`:
		i, e = strconv.ParseBool(v)
	case `string`:
		i = v
	case `[]byte`:
		i = []byte(v)
	}

	switch t {
	case `int`:
		i = int(i.(int64))
	case `int8`:
		i = int8(i.(int64))
	case `int16`:
		i = int16(i.(int64))
	case `int32`:
		i = int32(i.(int64))
	case `int64`:
		i = i.(int64)

	case `uint`:
		i = uint(i.(uint64))
	case `uint8`:
		i = uint8(i.(uint64))
	case `uint16`:
		i = uint16(i.(uint64))
	case `uint32`:
		i = uint32(i.(uint64))
	case `uint64`:
		i = i.(uint64)

	case `rune`:
		i = rune(i.(int64))
	case `byte`:
		i = byte(i.(uint64))
	case `float32`:
		i = float32(i.(float64))
	}

	return
}

func InterfaceToFloat64(v interface{}) float64 {
	switch t := v.(type) {
	case int:
		return float64(t)
	case int8:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)

	case uint:
		return float64(t)
	case uint8:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)

	case float32:
		return float64(t)
	case float64:
		return t
	}

	return 0
}
