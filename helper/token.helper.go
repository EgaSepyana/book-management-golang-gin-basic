package helper

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SingedDetails struct {
	Email string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenereteToken(email string) (singedToken string, singedRefreshToken string, err error) {
	claims := &SingedDetails{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(3)).Unix(),
		},
	}

	refreshClaims := &SingedDetails{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(6)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Fatal(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Fatal(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken(singedToken string) (claims *SingedDetails, msg string) {
	token, err := jwt.ParseWithClaims(singedToken, &SingedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SingedDetails)
	if !ok {
		msg = "Token is Invalid"
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "Token is Expired"
		msg = err.Error()
		return
	}

	return claims, msg
}
