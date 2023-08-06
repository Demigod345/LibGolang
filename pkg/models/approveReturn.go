package models

import (
	"fmt"
)

func ApproveReturn(requestId int, bookId int) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	updateSql := `UPDATE books SET available = available + 1 WHERE bookId = ?;`
	_, err = db.Exec(updateSql, bookId)
	if err != nil {
		fmt.Printf("error %s updating the database", err)
	} else {
		fmt.Printf("successfully updated the database ")
	}

	update2Sql := `DELETE FROM requests WHERE requestId= ? AND state = 'checkedIn';`
	_, err = db.Exec(update2Sql, requestId)
	if err != nil {
		fmt.Printf("error %s updating the database", err)
	} else {
		fmt.Printf("successfully deleted request from the database ")
	}

	fmt.Println("Models ApproveReturn() Function")
}