package util

import (
	"net/http"
)

// This exists because we cant specify route methods when using Go's default serveMux
func MethodGuard(handlerFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			WriteError(rw, http.StatusForbidden, "FORBIDDEN")
			return
		}
		handlerFunc(rw, r)
	}
}
