package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int64
	Name  string
	Email string
	Role  UserRole
}

type UserRole int

const (
	Admin UserRole = iota
	Client
)

var DB *sql.DB
var db *sql.DB

func Conect() {

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "0.0.0.0:3306",
		DBName:               "kevdb",
		AllowNativePasswords: true,
	}

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		// panic(pingErr)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Connected!")

	DB = db
}

func GetUsers() ([]User, error) {
	// An users slice to hold data from returned rows.
	var users []User

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, fmt.Errorf("getUsers %q: %v", "", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var usr User

		if err := rows.Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Role); err != nil {
			return nil, fmt.Errorf("getUsers %q: %v", "", err)
		}
		users = append(users, usr)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getUsers %q: %v", "", err)
	}
	return users, nil
}
