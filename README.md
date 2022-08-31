<p align="center">
  <h1 align="center">GoH - Go Helpers package to help with golang code creation productivity</h1>

<p align="center">
    <a href="https://pkg.go.dev/github.com/joaocprofile/goh"><img alt="Reference" src="https://img.shields.io/badge/go-reference-purple?style=for-the-badge"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-red.svg?style=for-the-badge"></a>
  </p>
</p>
<br>

This project was created for study purposes and use in personal projects, feel free to give me any feedback about the project, improvements, add new features, structure or anything you can report that can improve the project and spread knowledge!

<br>

## Functionalities:

- <b>(core)</b> Some useful functions for shared use, convert, encoder, types etc.
- <b>(database)</b> singleton's for database connection management and cache, SQL DML generator to facilitate database creation and manipulation
- <b>(environment)</b> Helpers for reading and configuring environment variables (.env) we can have different configurations for each environment (production, development).
- <b>(httprest)</b> Helpers for creating and configuring the Server, Routes, Database initialization and caching...
- <b>(httpwr)</b> Helpers to help with handling and handling Request and Response, Json, Errors, Body Reading, Params and Querys...

<br>

## Get started

### Installation

```sh
go get github.com/joaocprofile/goh
```

```sh
Create the .env file in your project ./
```

<br>

## Examples:

### Creation of Controllers

```go
// user_controller.go

package users

import (
	"net/http"
	"strconv"

	"github.com/joaocprofile/goh/httpwr"
)

type User struct {
	Id   uint `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{Id: 1, Name: "John doe"},
	{Id: 2, Name: "Big john"},
}

func Get(w http.ResponseWriter, r *http.Request) {
	sid, err := httpwr.Params("id", w, r)
	if err != nil {
		return
	}

	id, _ := strconv.Atoi(sid)
	id -= 1
	user := users[id]

	httpwr.Response(w, http.StatusOK, user)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	httpwr.Response(w, http.StatusOK, users)
}
```

### Creation of routes

```go
// users_routes.go

import (
	"net/http"

	"github.com/joaocprofile/goh/core"
)

var usersRoutes = []core.Route{
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
```

```go
// main.go
// Simple HTTP server

import (
	"net/http"

	"github.com/joaocprofile/goh/environment"
	"github.com/joaocprofile/goh/httprest"
)

func main() {
	environment.Inicialize()
	apiRest := httprest.NewHttpRest()
	apiRest.AddRoutes(usersRoutes)
	apiRest.ListenAndServe()
}
```

<br>

## ✨ Used packages

The following packages were directly used in this project::

- [Mux](https://github.com/gorilla/mux)
- [godotenv](github.com/joho/godotenv)
- [JSON Web Token](github.com/lib/pq)
- [Go postgres driver](github.com/lib/pq)
- [Redis](github.com/go-redis/redis/v8)
- [Gommon](github.com/labstack/gommon)

<br>

## My contacts:

- Email: joaocprofile@gmail.com
- Connect with me at [LinkedIn](https://www.linkedin.com/in/joaocprofile/).

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/joaocprofile/goh/LICENSE) file for details
