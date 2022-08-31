package httpwr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joaocprofile/goh/core"
)

func ReadBody(w http.ResponseWriter, r *http.Request, tStruct interface{}) error {
	dType := reflect.TypeOf(tStruct)
	if err := core.ShouldBeStruct(dType); err != nil {
		return err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(w, http.StatusUnprocessableEntity, errors.New("Error reading request body: "+err.Error()))
		return err
	}
	if err = json.Unmarshal(body, &tStruct); err != nil {
		Error(w, http.StatusInternalServerError, errors.New("JSON structure error: "+err.Error()))
		return err
	}

	return nil
}

func Params(param string, w http.ResponseWriter, r *http.Request) (string, error) {
	params := mux.Vars(r)
	id := params[param]
	if id == "" {
		err := errors.New("Error reading parameter " + param)
		Error(w, http.StatusBadRequest, err)
		return "", err
	}
	return id, nil
}

func Query(filter string, r *http.Request) string {
	return strings.ToLower(r.URL.Query().Get(filter))
}

func QueryToStruct(r *http.Request, queryStruct interface{}) error {
	objStruct := reflect.TypeOf(queryStruct)
	if err := core.ShouldBeStruct(objStruct); err != nil {
		return err
	}

	queryURLStrings := r.URL.Query()
	objValue := reflect.ValueOf(queryStruct)
	for iQueryParam, vQueryParam := range queryURLStrings {
		for iStructField := 0; iStructField < objStruct.Elem().NumField(); iStructField++ {
			field := objStruct.Elem().Field(iStructField)
			value := objValue.Elem().Field(iStructField)
			if strings.ToLower(field.Name) == strings.ToLower(iQueryParam) {
				value.SetString(vQueryParam[0])
			}
		}
	}

	return nil
}
