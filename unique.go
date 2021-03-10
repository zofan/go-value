package value

import (
	"fmt"
	"strings"
)

func UniqueStrings(list []string) (uniq []string) {
	keys := make(map[string]struct{})

	for _, v := range list {
		if _, ok := keys[v]; !ok {
			keys[v] = struct{}{}
			uniq = append(uniq, v)
		}
	}

	return
}

func UniqueGeneralStrings(list []string) (uniq []string) {
	keys := make(map[string]struct{})

	for _, v := range list {
		k := strings.ToLower(v)
		if _, ok := keys[k]; !ok {
			keys[k] = struct{}{}
			uniq = append(uniq, v)
		}
	}

	return
}

func UniqueInterfaces(list []interface{}) (uniq []interface{}) {
	keys := make(map[string]struct{})

	for _, v := range list {
		k := fmt.Sprintf(`%#v`, v)
		if _, ok := keys[k]; !ok {
			keys[k] = struct{}{}
			uniq = append(uniq, v)
		}
	}

	return
}
