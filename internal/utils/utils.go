package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)




func HashedPassword(pass string) (string,error){
	bytes,err:=bcrypt.GenerateFromPassword([]byte(pass),10)
	return string(bytes),err
}
func CheckPass(pass,hash string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(pass))
	return err==nil
}
func GenrateToken(len int) string{
	bys:=make([]byte, len)
	if _,err:=rand.Read(bys);err!=nil{
		log.Fatal("Errror in token generation")
	}
	return base64.URLEncoding.EncodeToString(bys)
}