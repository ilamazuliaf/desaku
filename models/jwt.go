package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	APP_NAME                  = "Rest Api Desa Kambingan Barat"
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY         = []byte("Apasaja Boleh :p")
)
