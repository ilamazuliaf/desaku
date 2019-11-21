package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ilamazuliaf/desaku/models"

	jwt "github.com/dgrijalva/jwt-go"
)

type M map[string]interface{}

// var APP_NAME = "SIMPLE APP JWT"
// var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
// var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
// var JWT_SIGNATURE_KEY = []byte("ikehikehkimochi")

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		tokenToString := r.Header.Get("x-token")
		token, err := jwt.Parse(tokenToString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != models.JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}
			return models.JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(M{
				"message": err.Error(),
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(M{
				"message": err.Error(),
			})
			return
		}
		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
