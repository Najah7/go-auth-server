package db

// use sqlite3
import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Najah7/go-auth-server/models"
)

// dbDriver is a database driver
const dbDriver = "sqlite3"

// dbSource is a database source
const dbSource = "./database.db"

// DB is a database
type DB struct {
	*sql.DB
}

// OpenDB opens a database
func OpenDB() (*DB, error) {
	// open database
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}, nil
}

// InitDB initializes a database
func InitDB(db *DB) (sql.Result, error) {
	// create table
	sqlStmt := `
	create table if not exists users (
		user_id text not null primary key,
		password text,
		nickname text,
		comment text,
		created_at datetime,
		updated_at datetime
	);
	`
	result, err := db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("db.Exec failed: %v", err)
	}
	return result, nil
}

// InsertUser inserts a user
func InsertUser(db *DB, user_id string, password string, nickname string, comment string) (models.User, error) {
	// insert user
	stmt, err := db.Prepare("insert into users(user_id, password, nickname, comment, created_at, updated_at) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return models.User{}, err
	}
	_, err = stmt.Exec(user_id, password, nickname, comment, time.Now(), time.Now())
	if err != nil {
		return models.User{}, err
	}
	return SelectUser(db, user_id)
}

// SelectUser selects a user
func SelectUser(db *DB, user_id string) (models.User, error) {
	// select user
	user := models.User{}
	err := db.QueryRow("select * from users where user_id = ?", user_id).Scan(&user.UserID, &user.Password, &user.Nickname, &user.Comment, &user.Created, &user.Updated)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUser updates a user
func UpdateUser(db *DB, user_id string, password string, nickname string, comment string) (models.User, error) {
	// update user
	stmt, err := db.Prepare("update users set password = ?, nickname = ?, comment = ?, updated_at = ? where user_id = ?")
	if err != nil {
		return models.User{}, err
	}
	_, err = stmt.Exec(password, nickname, comment, time.Now(), user_id)
	if err != nil {
		return models.User{}, err
	}
	return SelectUser(db, user_id)
}

// DeleteUser deletes a user
func DeleteUser(db *DB, user_id string) error {
	// delete user
	stmt, err := db.Prepare("delete from users where user_id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user_id)
	if err != nil {
		return err
	}
	return nil
}
