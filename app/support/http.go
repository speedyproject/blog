package support

import "net/http"

func GetRequestIP(req *http.Request) string {

	lastId := req.Header.Get("x-forwarded-for")

	if lastId == "" || lastId == "unknown" {
		lastId = req.Header.Get("Proxy-Client-IP")
	}
	if lastId == "" || lastId == "unknown" {
		lastId = req.Header.Get("WL-Proxy-Client-IP")
	}

	return lastId
}
