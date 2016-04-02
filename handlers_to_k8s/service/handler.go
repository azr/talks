package main

//START OMIT
import (
	"encoding/json"
	"net/http"
	"strconv"
)

func init() { http.HandleFunc("/user", UserGetHandler) }

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("name")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, found := DB.UserGet(id)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&user)
}

//END OMIT

var DB = FakeDB{}

type FakeDB map[uint64]User

type User struct {
	ID   uint64
	Name string
}

func (f FakeDB) UserGet(id uint64) (User, bool) {
	user, found := f[id]
	return user, found
}

// func main() {
// 	DB[0] = User{
// 		ID:   0,
// 		Name: "first !",
// 	}
// 	http.ListenAndServe(":8080", nil)
// }
