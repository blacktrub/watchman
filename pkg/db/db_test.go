package db

import (
	"fmt"
	"os"
	"testing"
)

func TestPrepareDBMustCreateTables(t *testing.T) {
	db, _ := GetDB()
	db.PrepareDB()

	tables := []string{"user", "project"}
	for i := 0; i < len(tables); i++ {
		table := tables[i]
		row := db.oneRow("select name from sqlite_master where name = ?", table)
		var dbTableName string
		err := row.Scan(&dbTableName)
		if err != nil {
			t.Error("Error happened", err)
		}

		if dbTableName != table {
			t.Error(fmt.Sprintf("Table %s not found", table))
		}
	}
	os.Remove(databasePath)
}

func TestGetUser(t *testing.T) {
	db, _ := GetDB()
	db.PrepareDB()

	user := User{1, 1}
	db.exec("insert into user (id, telegram_id) values (?, ?)", user.id, user.telegram_id)

	test_user, err := db.GetUser(1)
	if err != nil {
		t.Error("GetUser returned error", err)
	}

	if test_user.id != user.id {
		t.Error("test_user id is not equal", test_user.id, user.id)
	}
	os.Remove(databasePath)
}
