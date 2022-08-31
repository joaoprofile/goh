package users

import (
	"net/http"
	"strconv"

	"github.com/joaocprofile/goh/httpwr"
)

type User struct {
	Id   uint   `json:"id"`
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
