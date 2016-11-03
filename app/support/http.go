package support

import (
	"github.com/revel/revel"
)

func GetRequestIP(req *revel.Request) string {

	lastId := req.Header.Get("x-forwarded-for")

	if lastId == "" || lastId == "unknown" {
		lastId = req.Header.Get("Proxy-Client-IP")
	}
	if lastId == "" || lastId == "unknown" {
		lastId = req.Header.Get("WL-Proxy-Client-IP")
	}

	return lastId
}
