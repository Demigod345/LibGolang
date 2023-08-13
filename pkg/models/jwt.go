package models

import (
	"os"
	"gopkg.in/yaml.v3"
)

type JWTSecretKey struct {
	JWTSecretKey string `yaml:"JWTSecretKey"`
}

func GetJWTSecretKey() (string, error) {
	file, err := os.ReadFile("db.yaml")
	if err != nil {
		return "", err
	}
	var key JWTSecretKey
	if err := yaml.Unmarshal(file, &key); err != nil {
		return "", err
	}
	return key.JWTSecretKey, nil
}