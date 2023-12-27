package db

import (
	"testing"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// RemoveFile removes a file
func removeFile(filename string) {
    err := os.Remove(filename)
    if err != nil {
		panic(err)
	}
}

// DBファイルを削除するテストラッパー関数
func TestMain(m *testing.M) {
	// before test
	
	// run test
    exitVal := m.Run()

	// after test
    removeFile(dbSource)
    os.Exit(exitVal)
}

func TestOpenDB(t *testing.T) {
	// open database
	db, err := OpenDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// check database
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
}


func TestInitDB(t *testing.T) {
	db, err := OpenDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	InitDB(db)

	// check if the table users exists
	_, err = db.Query("SELECT * FROM users")
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertUser(t *testing.T) {
	// open database
	db, err := OpenDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// insert user
	InsertUser(db, "test_id", "test_password", "test_nickname", "test_comment")
	user, err := SelectUser(db, "test_id")
	if err != nil {
		t.Fatal(err)
	}

	// check if the user is inserted
	if user.UserID != "test_id" {
		t.Fatal("user.UserID != test_id")
	}
	if user.Password != "test_password" {
		t.Fatal("user.Password != test_password")
	}
	if user.Nickname != "test_nickname" {
		t.Fatal("user.Nickname != test_nickname")
	}
	if user.Comment != "test_comment" {
		t.Fatal("user.Comment != test_comment")
	}
}

func TestSelectUser(t *testing.T) {
	db, err := OpenDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// insert user
	InsertUser(db, "test_id", "test_password", "test_nickname", "test_comment")
	
	// select user
	user, err := SelectUser(db, "test_id")
	if err != nil {
		t.Fatal(err)
	}

	// check if the user is selected
	if user.UserID != "test_id" {
		t.Fatal("user.UserID != test_id")
	}
}

func TestUpdateUser(t *testing.T) {
	// open database
	db, err := OpenDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// insert user
	InsertUser(db, "test_id", "test_password", "test_nickname", "test_comment")

	// update user
	UpdateUser(db, "test_id", "new_password", "new_nickname", "new_comment")
	
	// check if the user is updated
	user, err := SelectUser(db, "test_id")
	if err != nil {
		t.Fatal(err)
	}

	// check if the user is updated
	if user.UserID != "test_id" {
		t.Fatal("user.UserID != test_id")
	}
	if user.Password != "new_password" {
		t.Fatal("user.Password != new_password")
	}
	if user.Nickname != "new_nickname" {
		t.Fatal("user.Nickname != new_nickname")
	}
	if user.Comment != "new_comment" {
		t.Fatal("user.Comment != new_comment")
	}
}
