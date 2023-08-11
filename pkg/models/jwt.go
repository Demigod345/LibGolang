package models

import (
	"log"
	"os"
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