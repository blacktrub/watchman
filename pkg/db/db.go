package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var databasePath = "./storage.db"

type User struct {
	id          int
	telegram_id int
}

type Project struct {
	id      int
	hash    string
	user_id int
}

type Database struct {
	d *sql.DB
}

func GetDB() (Database, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return Database{db}, err
	}
	return Database{db}, nil
}

func (db *Database) Close() error {
	return db.d.Close()
}

func (db *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.exec(query, args...)
}

func (db *Database) exec(query string, args ...interface{}) (sql.Result, error) {
	return db.d.Exec(query, args...)
}

func (db *Database) oneRow(query string, args ...interface{}) *sql.Row {
	return db.d.QueryRow(query, args...)
}

func (db *Database) PrepareDB() error {
	_, err := db.exec(`
	CREATE TABLE IF NOT EXISTS
	user(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    telegram_id INTEGER NOT NULL
	)`)
	if err != nil {
		return err
	}

	_, err = db.exec(`
	CREATE TABLE IF NOT EXISTS
	project(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    hash VARCHAR(32) NOT NULL,
	    user_id INTEGER NOT NULL,
	        FOREIGN KEY(user_id) REFERENCES user(id)
	)`)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetUser(id int) (User, error) {
	var user User
	row := db.oneRow("select id, telegram_id from user where id = ?", id)
	if err := row.Scan(&user.id, &user.telegram_id); err != nil {
		return user, err
	}
	return user, nil
}

