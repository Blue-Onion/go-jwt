package hanlder

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Blue-Onion/go-jwt/internal/utils"
	"github.com/Blue-Onion/go-jwt/store"
)



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
		_, ok := store.Users[username]
		if ok {
			err := http.StatusNotAcceptable
			http.Error(w, "User already exists", err)
			return
		}
		HashedPass,_:=utils.HashedPassword(password)
		store.Users[username]=store.LoginStruct{
			HashedPass: HashedPass,
		}
		fmt.Println(&w,"User Registered")
		fmt.Println()
	}
}
func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := http.StatusMethodNotAllowed
			http.Error(w, "Invalid Method", err)
			return
		}
		username := r.FormValue("name")
		password := r.FormValue("password")
		user,ok:=store.Users[username]
		if !ok{
			err := http.StatusUnauthorized
			http.Error(w, "User doesnt exists", err)
			return
		}
		isAuthorized:=utils.CheckPass(password,user.HashedPass)
		if !isAuthorized{
			err := http.StatusUnauthorized
			http.Error(w, "Password doesnt match", err)
			return

		}
		sessionToken:=utils.GenrateToken(32)
		csrfToken:=utils.GenrateToken(32)
		http.SetCookie(w,&http.Cookie{
			Name: "seesion_token",
			Value: sessionToken,
			Expires: time.Now().Add(7* time.Hour),
			HttpOnly: true,
			Secure: true,
		})
		http.SetCookie(w,&http.Cookie{
			Name: "csrf_token",
			Value: csrfToken,
			Expires: time.Now().Add(7* time.Hour),
			HttpOnly: false,
			Secure: false,
		})
		user.SessionToken=sessionToken
		user.CSRFToken=csrfToken
		store.Users[username]=user
		w.Write([]byte(user.SessionToken))
	}
}
func ProtectedRoute() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method!=http.MethodPost{
			http.Error(w,"Invalid Method",http.StatusMethodNotAllowed)
			return
		}
		if err:=utils.Authorrized(r);err!=nil{
			fmt.Println(err.Error())
			http.Error(w,"Unauthorized",http.StatusUnauthorized)
			return
		}
		username:=r.FormValue("name")
		w.Write([]byte(username))
	}
}
func LogOut() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		if err:=utils.Authorrized(r);err!=nil{
			http.Error(w,"Unauthorized",http.StatusUnauthorized)
			return
		} 
		http.SetCookie(w,&http.Cookie{
			Name: "seesion_token",
			Value: "",
			Expires: time.Now().Add(7* time.Hour),
			HttpOnly: true,
			Secure: true,
		})
		http.SetCookie(w,&http.Cookie{
			Name: "csrf_token",
			Value: "",
			Expires: time.Now().Add(-7* time.Hour),
			HttpOnly: false,
			Secure: false,
		})
		username:=r.FormValue("name")
		user,_:=store.Users[username]
		user.CSRFToken=""
		user.SessionToken=""
		store.Users[username]=user
		w.Write([]byte(username))

	}
}