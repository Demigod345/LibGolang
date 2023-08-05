package types

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

type CompleteRequest struct {
	RequestId int
	BookId int		`json:"bookId"`
	UserId int		`json:"userId"`
	State string
	CreatedAt string
	UpdatedAt string
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	PasswordC string `json:"passwordC"`
}

type DBInfo struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
}
