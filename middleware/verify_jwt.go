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
		if(auth==""){
			w.WriteHeader(http.StatusUnauthorized)
			res.Message = "Authorization header is empty"
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				log.Fatalln(err)
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
			log.Fatalln(err)
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			res.Message = "Your token is invalid"
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				log.Fatalln(err)
			}
		}

		next.ServeHTTP(w, r)
	})
}
