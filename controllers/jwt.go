package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var (
	key      []byte
	jwtToken *jwt.Token
	jwtStr   string
	signErr	 error
)

func viperEnvConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("There is an error with loading .env key value", err)
	}
}

func CreateJwt(email string, name string) string {
	viperEnvConfig()

	strKey := fmt.Sprintf("%v", viper.Get("JWT_KEY"))
	key = []byte(strKey)
	fmt.Printf("key = %v", key)
	jwtToken = jwt.NewWithClaims(jwt.SigningMethodHS256,
							jwt.MapClaims{
								"email": email,
								"name": name,
								"exp": time.Now().Add(time.Hour * 12).Unix(),
							})
	jwtStr, signErr = jwtToken.SignedString(key)

	if(signErr != nil){
		fmt.Printf("signErr = %v", signErr)
		return ""
	}

	return jwtStr
}