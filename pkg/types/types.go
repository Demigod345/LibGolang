package types

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string `json:"Username"`
	UserId   int    `json:"UserId"`
	IsAdmin  bool   `json:"IsAdmin"`
	jwt.RegisteredClaims
}

type User struct {
	UserId   int
	UserName string
	IsAdmin  bool
	Hash     string
}

type CompleteBook struct {
	BookId        int
	Title         string
	TotalQuantity int
	Available     int
}

type BookList struct {
	Books    []CompleteBook `json:"Title"`
	Message  interface{}    `json:"Message"`
	Username string         `json:"Username"`
}

type CompleteRequest struct {
	RequestId     int    `json:"RequestId"`
	BookId        int    `json:"BookId"`
	Title         string `json:"Title"`
	Available     int    `json:"Available"`
	TotalQuantity int    `json:"TotalQuantity"`
	UserId        int    `json:"UserId"`
	State         string
}

type RequestList struct {
	Requests []CompleteRequest
	Message  interface{} `json:"Message"`
	Username string
}

type DBInfo struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
}

type PageMessage struct {
	Message interface{} `json:"Message"`
}

type FileName struct {
	AdminHome           string
	UserHome            string
	Login               string
	PageNotFound        string
	UnauthorizedAccess  string
	InternalServerError string
	Signup              string
}
