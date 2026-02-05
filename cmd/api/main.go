package main

import (
	"net/http"

	hanlder "github.com/Blue-Onion/go-jwt/internal/handler"
)



func main(){
	http.HandleFunc("/register",hanlder.Register())
	http.HandleFunc("/login",hanlder.Login())
	http.HandleFunc("/logout",hanlder.LogOut())
	http.HandleFunc("/protected",hanlder.ProtectedRoute())
	http.ListenAndServe(":8080",nil)

}