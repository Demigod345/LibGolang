package models

import (
	"fmt"
)

func ApproveAdmin(requestId int, userId int) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	updateSql := `UPDATE users SET isAdmin = 1 WHERE userId = ?;`
	_, err = db.Exec(updateSql, userId)
	if err != nil {
		fmt.Printf("error %s updating the database", err)
	} else {
		fmt.Printf("successfully updated the database ")
	}

	update2Sql := `DELETE FROM requests WHERE requestId= ? AND state = 'AdminRequest';`
	_, err = db.Exec(update2Sql, requestId)
	if err != nil {
		fmt.Printf("error %s updating the database", err)
	} else {
		fmt.Printf("successfully updated the database ")
	}

	fmt.Println("Models ApproveAdmin() Function")
}
