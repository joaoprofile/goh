package main

import (
	"github.com/joaocprofile/goh/exemplos/simples/users"

	"github.com/joaocprofile/goh/environment"
	"github.com/joaocprofile/goh/httprest"
)

func main() {
	environment.Inicialize()
	apiRest := httprest.NewHttpRest()
	apiRest.AddRoutes(users.Routes)
	apiRest.ListenAndServe()
}
