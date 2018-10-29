package entity

const initUserTable = `CREATE TABLE IF NOT EXISTS "users" (
    "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
    "username" VARCHAR(64) UNIQUE NOT NULL,
	"password" CHAR(44) NOT NULL,
	"email" VARCHAR(50) UNIQUE NOT NULL,
    "createdTime" DATE NOT NULL
);`

const initMeetingTable = `CREATE TABLE IF NOT EXISTS "meetings"(
	"mid" INTEGER PRIMARY KEY AUTOINCREMENT,
	"title" VARCHAR(100) UNIQUE NOT NULL,
	"creatorid" INTEGER NOT NULL,
	"start" DATETIME NOT NULL,
	"end" DATETIME NOT NULL,
	FOREIGN KEY (creatorid) REFERENCES users(uid)
);`

const initParticipantTable = `CREATE TABLE IF NOT EXISTS "participants"(
	"uid" INTEGER,
	"mid" INTEGER,
	FOREIGN KEY(uid) REFERENCES users(uid),
	FOREIGN KEY(mid) REFERENCES meetings(mid)
);`

const addUser = `INSERT INTO users(username, password, email, createdTime) values(?, ?, ?, ?)`

const deleteUser = `DELETE FROM users WHERE username=?`

const getUserByName = `SELECT * FROM users WHERE username=?`

const getUserByEmail = `SELECT * FROM users WHERE email=?`

const getAllUser = `SELECT * FROM users`

const addMeeting = `INSERT INTO meetings(title, creatorid, start, end) values(?, ?, ?, ?)`

const getParticipating = `SELECT title,start,end FROM meetings WHERE mid IN (SELECT mid FROM participants WHERE uid=?)`

const getCreatedMeeting = `SELECT title,start,end FROM meetings WHERE creatorid=?`

const getMeetingByTitle = `SELECT * FROM meetings WHERE title=?`

const addParticipant = `INSERT INTO participants(uid, mid) values(?, ?)`


