package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

type SignedDetails struct {
	Username string
	ID       int
	jwt.RegisteredClaims
}

func CreateToken(username string, ID int) (string, string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &SignedDetails{
		Username: username,
		ID:       ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    username,
			ID:        strconv.Itoa(ID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	refreshClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(168))),
		},
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	refreshToken, err := refreshClaims.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		//msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token in invalid")
		//msg = err.Error()
		return
	}

	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		msg = fmt.Sprintf("token is expired")
		//msg = err.Error()
	}

	return claims, msg
}
