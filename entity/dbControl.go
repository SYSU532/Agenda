package entity

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var agendaDB *sql.DB
var addUserStmt, getUserByNameStmt, getUserByEmailStmt *sql.Stmt

const dbPath = "./data"
const dbFile = "agenda.db"

func init() {
	var err error

	//Create database directory if not exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		os.Mkdir(dbPath, os.ModePerm)
	}

	//Open DB
	agendaDB, err = sql.Open("sqlite3", filepath.Join(dbPath, dbFile))
	checkErr(err)

	//Init table if not exist.
	_, err = agendaDB.Exec(initUserTable)
	checkErr(err)
	_, err = agendaDB.Exec(initMeetingTable)
	checkErr(err)

	//Prepare statements
	addUserStmt, err = agendaDB.Prepare(addUser)
	checkErr(err)
	getUserByNameStmt, err = agendaDB.Prepare(getUserByName)
	checkErr(err)
	getUserByEmailStmt, err = agendaDB.Prepare(getUserByEmail)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func LoginUser(username, password string) error {
	result, err := getUserByNameStmt.Query(username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	if result.Next() {
		passHash := sha256.Sum256([]byte(password))
		hashStr := base64.URLEncoding.EncodeToString(passHash[:])
		var uid int
		var uname, pass, email string
		var ctime time.Time
		err = result.Scan(&uid, &uname, &pass, &email, &ctime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return errors.New("database scan error")
		}
		if pass == hashStr {
			return nil
		} else {
			return errors.New("wrong password")
		}
	} else {
		return errors.New("username does not exist")
	}
}

func AddUser(username, password, email string) error {
	err := checkUserDuplicate(username, email)
	if err != nil {
		return err
	}
	passHash := sha256.Sum256([]byte(password))
	hashStr := base64.URLEncoding.EncodeToString(passHash[:])
	datetime := time.Now().Format(time.RFC3339)
	_, err = addUserStmt.Exec(username, hashStr, email, datetime)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return errors.New("database error when adding user")
	}
	return nil
}

func checkUserDuplicate(username, email string) error {
	result, err := getUserByNameStmt.Query(username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	if result.Next() {
		return errors.New("username already exists")
	}
	result, err = getUserByEmailStmt.Query(email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	if result.Next() {
		return errors.New("email already exists")
	}
	return nil
}
