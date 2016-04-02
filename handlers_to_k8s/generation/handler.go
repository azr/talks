package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"gopkg.in/validator.v2"
)

//START_REQ OMIT
type UserCreateRequest struct {
	Name string `validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
}

//END_REQ OMIT

func HTTPUserCreateRequest(r *http.Request) (uc UserCreateRequest, err error) {
	err = json.NewDecoder(r.Body).Decode(&uc)
	if err != nil {
		return
	}
	err = validator.Validate(uc)
	return
}

//BEFORE_USER_CREATE_GENERATE OMIT

//go:generate varhandler -func UserCreate
//BEFORE_USER_CREATE OMIT
func UserCreate(u UserCreateRequest) (status int, err error) {
	//USER_CREATE_FIRST_LINE OMIT
	created := DB.CreateUser(u)
	if !created {
		return 0, errors.New("username taken")
	}
	return http.StatusCreated, nil
}

//AFTER_USER_CREATE OMIT

//go:generate varhandler -func UserGet
func UserGet(un UserName) (u User, status int, err error) {
	//USER_GET_FIRST_LIST OMIT
	var found bool
	u, found = DB.GetUser(un)
	if !found {
		status = http.StatusNotFound
	}
	return
}

type UserName string

func HTTPUserName(r *http.Request) (UserName, error) {
	un := UserName(r.URL.Query().Get("name"))
	// validate ...
	return un, nil
}

//BEFORE_INIT OMIT
func init() {
	http.HandleFunc("/user/create", UserCreateHandler)
	http.HandleFunc("/user/get", UserCreateHandler)
}

//AFTER_INIT OMIT

var DB = FakeDB{}

type FakeDB map[string]User

type User struct {
	ID   uint64
	Name string
}

func (f FakeDB) CreateUser(u UserCreateRequest) (created bool) {
	_, found := f[u.Name]
	if found {
		return false
	}
	f[u.Name] = User{
		ID:   uint64(len(f) + 1),
		Name: u.Name,
	}
	return true
}

func (f FakeDB) GetUser(id UserName) (user User, found bool) {
	user, found = f[string(id)]
	return
}

func main() {
	http.ListenAndServe(":8080", nil)
}
