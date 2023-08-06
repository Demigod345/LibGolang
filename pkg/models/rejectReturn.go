package models

import (
	"fmt"
)

func RejectReturn(requestId int) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	updateSql := `UPDATE requests SET state = 'issued' WHERE requestId= ? AND state = 'checkedIn';`
	_, err = db.Exec(updateSql, requestId)
	if err != nil {
		fmt.Printf("error %s updating the database", err)
	} else {
		fmt.Printf("successfully updated the database ")
	}
	fmt.Println("Models RejectReturn() Function")
}