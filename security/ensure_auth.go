package security

import (
	"net/http"

	"github.com/joaocprofile/goh/httpwr"
)

func EnsureAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := TokenValidate(r); err != nil {
			httpwr.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
