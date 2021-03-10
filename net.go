package value

import (
	"golang.org/x/net/idna"
	"net/url"
	"strings"
)

func Hostname(v string) string {
	v = strings.ToLower(v)

	if strings.HasPrefix(v, `//`) {
		v = `http:` + v
	}

	if !strings.Contains(v, `://`) {
		v = `http://` + v
	}

	if u, err := url.Parse(v); err == nil {
		v = u.Hostname()
		v = strings.TrimPrefix(v, `www.`)
		v, err = idna.ToASCII(v)
		if err != nil {
			return ``
		}
	} else {
		return ``
	}

	return v
}
