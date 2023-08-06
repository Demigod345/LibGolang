package models

import (
	"fmt"
)

func ApproveRequest(requestId int, bookId int) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	updateSql := `UPDATE books SET available = available -1 WHERE bookId = ?;`
	_, err = db.Exec(updateSql, bookId)
	if err != nil {
		fmt.Printf("error %s updating the database", err)
	} else {
		fmt.Printf("successfully updated the database ")
	}

	update2Sql := `UPDATE requests SET state = 'issued' WHERE requestId= ? AND state = 'requested';`
	_, err = db.Exec(update2Sql, requestId)
	if err != nil {
		fmt.Printf("error %s updating the database", err)
	} else {
		fmt.Printf("successfully updated the database ")
	}

	fmt.Println("Models ApproveRequest() Function")
}
