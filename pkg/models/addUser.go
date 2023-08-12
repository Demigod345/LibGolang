package models


func AddUser(username string, hash string) error {
	db, err := Connection()
	if err != nil {
		return err
	}

	insertSql := "INSERT INTO users (userName, hash, isAdmin) VALUES(?,?, false)"
	_, err = db.Exec(insertSql, username, hash)
	if err != nil {
		return err
	} else {
		return nil
	}
}
