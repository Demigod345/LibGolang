package models

import ()

func AddUser(username string, hash string) error {
	db, err := Connection()
	if err != nil {
		return err
	}

	isAdmin:= false;

	insertSql := "INSERT INTO users (userName, hash, isAdmin) VALUES (?,?,?)"
	_, err = db.Exec(insertSql, username, hash, isAdmin)
	if err != nil {
		return err
	} else {
		return nil
	}
}