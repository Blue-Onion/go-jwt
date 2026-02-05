package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Blue-Onion/go-jwt/store"
)

var AuthErr=errors.New("Unauthorrized")
func Authorrized(r *http.Request) error{
	username:=r.FormValue("name")

	user,ok:=store.Users[username]
	fmt.Printf("User: %+v\n", user)
	if !ok{
		return AuthErr
	}

	csrf,err:=r.Cookie("csrf_token")
	if err!=nil{
		return AuthErr
	}

	if csrf.Value!=user.CSRFToken||csrf.Value==""{
		return AuthErr
	}
	return nil
}