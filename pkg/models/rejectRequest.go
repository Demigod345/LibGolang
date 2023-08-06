package models

import (
	"fmt"
)

func RejectRequest(requestId int) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	updateSql := `DELETE FROM requests WHERE requestId= ? AND state = 'requested';`
	_, err = db.Exec(updateSql, requestId)
	if err != nil {
		fmt.Printf("error %s updating the database", err)
	} else {
		fmt.Printf("successfully updated the database ")
	}

	fmt.Println("Models RejectRequest() Function")
}