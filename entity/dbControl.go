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
var addUserStmt, deleteUserStmt, getUserByNameStmt, getUserByEmailStmt, getAllUserStmt *sql.Stmt

const dbPath = "./data"
const dbFile = "agenda.db"

// Print finding user info structure
type FindUserInfo struct {
	Username string
	Email    string
}

func init() {
	var err error

	//Create database directory if not exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		os.Mkdir(dbPath, os.ModePerm)
	}

	//Open DB
	agendaDB, err = sql.Open("sqlite3", filepath.Join(dbPath, dbFile))
	checkErr(err)

	//Solve "db is locked" error
	agendaDB.Exec("PRAGMA journal_mode=WAL;")

	//Init table if not exist.
	_, err = agendaDB.Exec(initUserTable)
	checkErr(err)
	_, err = agendaDB.Exec(initMeetingTable)
	checkErr(err)
	_, err = agendaDB.Exec(initParticipantTable)
	checkErr(err)

	//Prepare statements
	addUserStmt, err = agendaDB.Prepare(addUser)
	checkErr(err)
	deleteUserStmt, err = agendaDB.Prepare(deleteUser)
	checkErr(err)
	getUserByNameStmt, err = agendaDB.Prepare(getUserByName)
	checkErr(err)
	getUserByEmailStmt, err = agendaDB.Prepare(getUserByEmail)
	checkErr(err)
	getAllUserStmt, err = agendaDB.Prepare(getAllUser)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Login user with username and password
func LoginUser(username, password string) error {
	result, err := getUserByNameStmt.Query(username)
	defer result.Close()
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

// Create user with username and password
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

// Check user creation is duplicate or not
func checkUserDuplicate(username, email string) error {
	result, err := getUserByNameStmt.Query(username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	if result.Next() {
		return errors.New("username already exists")
	}
	result.Close()
	result, err = getUserByEmailStmt.Query(email)
	defer result.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	if result.Next() {
		return errors.New("email already exists")
	}
	return nil
}

// Delete current user (Only be invoked in login state)
func DeleteUser() (string, error) {
	curUser, err := GetCurrentUser()
	if err != nil || curUser.Username == "" {
		return "", errors.New("not in Login state")
	}
	_, err = deleteUserStmt.Exec(curUser.Username)
	if err != nil {
		return "", err
	}
	return curUser.Username, nil
}

// Get all users with special conditions
func GetUserList(username string, email string) ([]FindUserInfo, error) {
	result, err := getAllUserStmt.Query()
	defer result.Close()
	var (
		output               []FindUserInfo
		uid                  int
		uname, upass, uemail string
		utime                time.Time
	)
	// Start checking with conditions
	if username == "" {
		if email != "" {
			result, err = getUserByEmailStmt.Query(email)
		}
	} else {
		result, err = getUserByNameStmt.Query(username)
	}
	for result.Next() {
		result.Scan(&uid, &uname, &upass, &uemail, &utime)
		output = append(output, FindUserInfo{Username: uname, Email: uemail})
	}
	return output, err
}

func Close() {
	agendaDB.Close()
}
