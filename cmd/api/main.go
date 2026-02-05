package main

import (
	"net/http"

	hanlder "github.com/Blue-Onion/go-jwt/internal/handler"
)



func main(){
	http.HandleFunc("/register",hanlder.Register())
	http.HandleFunc("/login",func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/logout",func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/protected",func(w http.ResponseWriter, r *http.Request) {})
	http.ListenAndServe(":8080",nil)

}