package http_testscenario

import "net/http"

type AuthProvider interface {
	AddAuth(req *http.Request)
}
