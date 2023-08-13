package types

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string `json:"username"`
	UserId   int    `json:"userId"`
	IsAdmin  bool   `json:"isAdmin"`
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
	Books    []CompleteBook `json:"title"`
	Message  interface{}    `json:"message"`
	Username string         `json:"username"`
}

type CompleteRequest struct {
	RequestId     int    `json:"requestId"`
	BookId        int    `json:"bookId"`
	Title         string `json:"title"`
	Available     int    `json:"available"`
	TotalQuantity int    `json:"totalQuantity"`
	UserId        int    `json:"userId"`
	State         string
}

type RequestList struct {
	Requests []CompleteRequest
	Message  interface{} `json:"message"`
	Username string
}

type DBInfo struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
}

type PageMessage struct {
	Message interface{} `json:"message"`
}
