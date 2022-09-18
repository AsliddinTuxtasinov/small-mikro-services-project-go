package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var MySigningKey = []byte("uwhiuwhoqwWYUGw234323@$%^&")

func IsAuthorized(endpoint http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Invalid signing method")
				}
				aud := "asliddin.jwtgo.io"
				if checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false); !checkAudience {
					return nil, fmt.Errorf("Invalid aud")
				}

				// iss := "jwt.io"
				// if checkIss := token.Claims.(jwt.MapClaims).VerifyAudience(iss, false); !checkIss {
				// 	return nil, fmt.Errorf("Invalid iss")
				// }

				return MySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "No authorization token provided")
		}
	})
}
