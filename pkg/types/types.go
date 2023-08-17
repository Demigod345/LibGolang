package types

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string 
	UserId   int    
	IsAdmin  bool   
	jwt.RegisteredClaims
}

type User struct {
	UserId   int
	UserName string
	IsAdmin  bool
	Hash     string
}

type RequestUser struct {
	UserName string
	Password string
	ConfirmPassword string
}

type CompleteBook struct {
	BookId        int
	Title         string
	TotalQuantity int
	Available     int
}

type BookList struct {
	Books    []CompleteBook 
	Message  interface{}    
	Username string         
}

type CompleteRequest struct {
	RequestId     int   
	BookId        int    
	Title         string 
	Available     int    
	TotalQuantity int    
	UserId        int    
	State         string
}

type RequestList struct {
	Requests []CompleteRequest
	Message  interface{} 
	Username string
}

type DBInfo struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
}

type PageMessage struct {
	Message interface{} 
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
