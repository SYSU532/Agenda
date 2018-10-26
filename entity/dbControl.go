package entity

import (
	"database/sql"
	
)

var userDB, meetingDB *sql.DB
var addUserStmt, getUserStmt *sql.Stmt

func init() {
	var err error

	//Open DB
	userDB, err = sql.Open("sqlite3", "./db/users.db")
	checkErr(err)
	meetingDB, err = sql.Open("sqlite3", "./db/meetings.db")
	checkErr(err)

	//Init table if not exist.
	_, err = userDB.Exec(initUserTable)
	checkErr(err)
	_, err = meetingDB.Exec(initMeetingTable)
	checkErr(err)

	//Prepare statements
	addUserStmt, err = userDB.Prepare(addUser)
	checkErr(err)
	getUserStmt, err = userDB.Prepare(getUser)
	checkErr(err)

	
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func LoginUser(username, password string) {

}

func AddUser(username, password, email string) {

}