package models

import (
	// "fmt"

	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error %v while creating mock database connection.", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"userId", "username", "hash", "isAdmin"}).AddRow(3, "qwer", "1", false)

	// AddRow(3, "Admin", "1", true).
	// AddRow(3, "myTestBook", "69", false).
	// AddRow(3, "yourTestBook", "490", false)

	username := "qwer6789"
	mock.ExpectQuery("SELECT \\* FROM users where username = ?").
		WithArgs(username).
		WillReturnRows(rows)

	// // fmt.Println(rows)

	fmt.Println(UserExists(db, "qwer6789"))
	result, _, err := UserExists(db, "qwer6789")
	// fmt.Println(bookList)
	if err != nil {
		t.Fatalf("Error %v getting FetchBooksDb() function", err)
	}

	// fmt.Println(user)
	expected := true

	if result != expected {
		t.Errorf("got %v, wanted %v", result, expected)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
