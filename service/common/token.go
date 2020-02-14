package common

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type Context struct{
	ID int64
	Username string
}

func SignJWT(c Context,secret string) (tokenstring string,err error){
	//Load the jwt secret from  config if the secret isn't specified.
	if(secret==""){
		secret = viper.GetString("jwt.secret")
	}

	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id" : c.ID,
		"username": c.Username,
		"nbf": time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	tokenstring,err = token.SignedString([]byte(secret))

	return tokenstring,err
}

//JWT鉴权