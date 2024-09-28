package middleware

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func viperEnvConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("There is an error with loading .env key value", err)
	}
}

func VerifyJwtToken(next http.Handler) http.Handler {
	type Response struct {
		Message string `json:"message"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		viperEnvConfig()
		var res Response

		auth := r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if(auth==""){
			w.WriteHeader(http.StatusUnauthorized)
			res.Message = "Authorization header is empty"
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				log.Fatalln(err)
			}
			return
		}
		if(len(strings.Split(auth, " "))!=2)||(strings.Split(auth, " ")[0]!="Bearer"){
			w.WriteHeader(http.StatusUnauthorized)
			res.Message = "Please follow 'Bearer yourJwtToken' format"
			err := json.NewEncoder(w).Encode(res)
			if(err!=nil){
				fmt.Println(err)
			}
			return
		}
		jwtToken := strings.Split(auth, " ")[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			strKey := fmt.Sprintf("%v", viper.Get("JWT_KEY"))
			secretKey := []byte(strKey)
			return secretKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			if !token.Valid {
				res.Message = "Your token is invalid"
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					fmt.Println(err)
				}
				return
			}
			res.Message = "There is error when parsing jwt token"
			err := json.NewEncoder(w).Encode(res)
			fmt.Println(err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
