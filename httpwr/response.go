package httpwr

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joaocprofile/goh/core"
)

type Erro struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statuscode int, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	if json != nil {
		if _, err := w.Write(json); err != nil {
			log.Println(core.Red(err.Error()))
		}
	}
}

func Response(w http.ResponseWriter, statuscode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println(core.Red(err.Error()))
		}
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	Response(w, statusCode, Erro{
		Code:  statusCode,
		Error: err.Error(),
	})
}
