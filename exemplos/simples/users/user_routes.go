package users

import (
	"net/http"

	"github.com/joaocprofile/goh/core"
)

var Routes = []core.Route{
	{
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Function:       Get,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Function:       GetAll,
		Authentication: false,
	},
}
