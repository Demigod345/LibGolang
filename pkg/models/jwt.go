package models

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"gopkg.in/yaml.v3"
)

type JWTSecretKey struct {
	JWTSecretKey string `yaml:"JWTSecretKey"`
}

func GetJWTSecretKey() string {
	file, err := os.ReadFile("db.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var key JWTSecretKey

	if err := yaml.Unmarshal(file, &key); err != nil {
		log.Fatal(err)
	}
	return key.JWTSecretKey

}

func GenerateJWT(userId int, username string, isAdmin bool) (string, error) {
	secretKey := []byte("hey")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["isAdmin"] = isAdmin
	claims["username"] = username
	claims["userId"] = userId

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	fmt.Println(GetJWTSecretKey())
	return tokenString, nil
}



