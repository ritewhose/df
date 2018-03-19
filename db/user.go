package db

import "fmt"

// UserDriver defines methods for interacting with User information.
type UserDriver interface {
	Exists(string) (bool, error)
	InsertUser(*User) error
	SelectUser(string) (*User, error)
}

// User represents a user. TODO:
// 		- add points
//		- logging?
type User struct {
	UserName string
}

// Exists checks whether a user exists in the database. Users are logged
// by discord ID, so nickname changes shouldn't cause any issues.
func (db *DB) Exists(userName string) (bool, error) {
	var exists bool

	row := db.QueryRow("SELECT EXISTS(SELECT * FROM User WHERE userName = ? LIMIT 1)", userName)
	err := row.Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("[Exists] %s", err)
	}

	return exists, nil
}

// InsertUser inserts a user into the database.
func (db *DB) InsertUser(u *User) error {
	statement, err := db.Prepare("INSERT INTO User(userName) VALUES (?)")
	if err != nil {
		return fmt.Errorf("[InsertUser] %s", err)
	}
	defer statement.Close()

	_, err = statement.Exec(u.UserName)
	if err != nil {
		return fmt.Errorf("[InsertUser] %s", err)
	}

	return nil
}

// SelectUser while not in use right now, I've included it just in case
// we need it later.
func (db *DB) SelectUser(userName string) (*User, error) {
	u := &User{}

	row := db.QueryRow("SELECT * FROM User WHERE userName = ? LIMIT 1", userName)
	err := row.Scan(&u)
	if err != nil {
		return nil, fmt.Errorf("[SelectUser] %s", err)
	}

	return u, nil
}
