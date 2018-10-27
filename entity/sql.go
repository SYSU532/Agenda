package entity

const initUserTable =
	`CREATE TABLE IF NOT EXISTS "users" (
    "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
    "username" VARCHAR(64) UNIQUE NOT NULL,
	"password" CHAR(44) NOT NULL,
	"email" VARCHAR(50) UNIQUE NOT NULL,
    "createdTime" DATE NOT NULL
);`

const initMeetingTable =
	`CREATE TABLE IF NOT EXISTS "meetings"(
	"mid" INTEGER PRIMARY KEY AUTOINCREMENT,
	"title" VARCHAR(100) NOT NULL,
	"creatorid" INTEGER NOT NULL,
	"time" DATE NOT NULL,
	FOREIGN KEY (creatorid) REFERENCES users(uid)
);`


const addUser =
	`INSERT INTO users(username, password, email, createdTime) values(?, ?, ?, ?)`

const getUserByName =
	`SELECT * FROM users WHERE username=?`

const getUserByEmail =
	`SELECT * FROM users WHERE email=?`





