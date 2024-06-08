package apitest

import "net/url"

func GetHostAndPortFromURL(u string) (host string, port string, err error) {
	uo, err := url.Parse(u)
	if err != nil {
		return
	}

	host = uo.Hostname()
	port = uo.Port()

	return
}
