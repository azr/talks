package main

//START OMIT
import (
	"encoding/json"
	"net/http"
	"strconv"
)

func init() { http.HandleFunc("/user", UserGetHandler) }

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, found, err := DB.UserGet(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(&user)
}

//END OMIT

var DB = FakeDB{}

type FakeDB map[uint64]User

type User struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (f FakeDB) UserGet(id uint64) (User, bool, error) {
	user, found := f[id]
	return user, found, nil
}

func main() {
	DB[1] = User{
		ID:   1,
		Name: "first !",
	}
	DB[2] = User{
		ID:   2,
		Name: "second !",
	}
	http.ListenAndServe(":8081", nil)
}
