package types

import "github.com/golang-jwt/jwt"

type Book struct {
	Title string		`json:"title"`
	Quantity string		`json:"quantity"`
}

type User struct {
	UserId int  
	UserName string
	IsAdmin bool
	Hash string
}

type CompleteBook struct {
	BookId int
	Title string
	TotalQuantity int
	Available int
}

type BookList struct {
	Books []CompleteBook `json:"title"`
}

type CompleteRequest struct {
	RequestId int	`json:"requestId"`
	BookId int		`json:"bookId"`
	UserId int		`json:"userId"`
	State string
	CreatedAt string
	UpdatedAt string
}

type RequestList struct {
	Requests []CompleteRequest
}

type AdminRequests struct {
	ApproveRequests RequestList
	IssuedBooks RequestList
	ReturnRequests RequestList
	AdminRequests RequestList
}

type SignupRequest struct {
	Username string 	`json:"username"`
	Password string 	`json:"password"`
	PasswordC string 	`json:"passwordC"`
}

type DBInfo struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
}

type Claims struct {
	jwt.StandardClaims
	Exp string  		`json:"exp"`
	Username string 	`json:"username"`
	IsAdmin bool 		`json:"isAdmin"`
	UserId int 			`json:"userId"`
}

