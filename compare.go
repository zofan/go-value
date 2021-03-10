package value

import "time"

func CmpDate(a, b time.Time) bool {
	return a.Truncate(time.Hour*24).Unix() == b.Truncate(time.Hour*24).Unix()
}

func CmpCrossStrings(a, b []string) bool {
	for _, av := range a {
		for _, bv := range b {
			if av == bv {
				return true
			}
		}
	}

	return false
}
