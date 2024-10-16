package config

import (
	"Dexun/form"
	"Dexun/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("secret")

type Claims struct {
	AccountName string
	jwt.StandardClaims
}

func ReleaseAgentCode(agent form.AgentInfo) (string, error) {
	expireTime := time.Now().Add(time.Minute * 1).Unix()
	claims := &Claims{
		AccountName: agent.TelNum,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "Raphael",
		},
	}
	code := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//codeString, err := code.SignedString(jwtKey)
	codetmp, err := code.SignedString(jwtKey)
	fmt.Println(codetmp)
	codeString := "12345"
	if err != nil {
		return "", err
	}
	return codeString, nil
}

func ReleaseCode(account form.AccountInfo) (string, error) {
	expireTime := time.Now().Add(time.Minute * 1).Unix()
	claims := &Claims{
		AccountName: account.TelNum,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "Raphael",
		},
	}
	code := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//codeString, err := code.SignedString(jwtKey)
	codetmp, err := code.SignedString(jwtKey)
	fmt.Println(codetmp)
	codeString := "12345"
	if err != nil {
		return "", err
	}
	return codeString, nil
}

func ReleaseToken(account model.Account) (string, error) {
	expireTime := time.Now().Add(time.Hour * 2).Unix()

	claims := &Claims{
		AccountName: account.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "Raphael",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
