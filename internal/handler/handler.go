package hanlder

import (
	"fmt"
	"net/http"

	"github.com/Blue-Onion/go-jwt/internal/utils"
)

type Login struct {
	HashedPass  string
	SessionToken string
	CSRFToken    string
}

var Users = map[string]Login{}

func Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			err := http.StatusMethodNotAllowed
			http.Error(w, "Invalid Method", err)
			return
		}
		username := r.FormValue("name")
		password := r.FormValue("password")
		if len(password) < 8 || len(username) == 0 {
			err := http.StatusNotAcceptable
			http.Error(w, "Invalid username/password", err)
			return
		}
		_, ok := Users[username]
		if ok {
			err := http.StatusNotAcceptable
			http.Error(w, "User already exists", err)
			return
		}
		HashedPass,_:=utils.HashedPassword(password)
		Users[username]=Login{
			HashedPass: HashedPass,
		}
		fmt.Println(&w,"User Registered")
		fmt.Println()
	}
}
