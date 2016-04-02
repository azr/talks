// Code generated by "varhandler -func UserCreate"; DO NOT EDIT

package main

import "net/http"

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	param0, err := HTTPUserCreateRequest(r)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusBadRequest, err)
		return
	}

	var status int

	status, err = UserCreate(param0)
	if err != nil {
		HandleHttpErrorWithDefaultStatus(w, r, http.StatusInternalServerError, err)
		return
	}

	if status != 0 {
		w.WriteHeader(status)
	}

}