package models

import ()

func ApproveAdmin(requestId int, userId int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	updateSql := `UPDATE users SET isAdmin = 1 WHERE userId = ?;`
	_, err = db.Exec(updateSql, userId)
	if err != nil {
		return "", err
	} 

	update2Sql := `DELETE FROM requests WHERE requestId= ? AND state = 'AdminRequest';`
	_, err = db.Exec(update2Sql, requestId)
	if err != nil {
		return "", err
	} 

	return "Successfully Approved Admin Request.", err
}
