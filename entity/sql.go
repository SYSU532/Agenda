package entity

const initUserTable =
	`CREATE TABLE IF NOT EXISTS "users" (
    "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
    "username" VARCHAR(64),
	"password" VARCHAR(20),
	"email" VARCHAR(50),
    "createdDate" DATE
);`

const initMeetingTable =
	`CREATE TABLE IF NOT EXISTS "meetings"(
	"mid" INTEGER PRIMARY KEY AUTOINCREMENT,
	"title" VARCHAR(100),
	"time" DATE
);`


const addUser =
	`INSERT INTO users(username, password, email) values(?, ?, ?)`

const getUser =
	`SELECT * FROM users WHERE username=?`


