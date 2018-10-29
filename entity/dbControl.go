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
var addUserStmt, deleteUserStmt, getUserByNameStmt, getUserByEmailStmt, getUserByIDStmt, getAllUserStmt *sql.Stmt
var addMeetingStmt, getPartStmt, getCreatedStmt, getTitleMeetingStmt, addPartStmt, getAllMeetingsStmt *sql.Stmt
var getAllParticipantsStmt *sql.Stmt

const dbPath = "./data"
const dbFile = "agenda.db"

// Print finding user info structure
type FindUserInfo struct {
	uid int
	Username string
	Email    string
}

// Meeting structure
type Meeting struct {
	StartTime string
	EndTime string
	Title string
	Creator string
	Participant []string
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
	getUserByIDStmt, err = agendaDB.Prepare(getUserNameByID)
	checkErr(err)
	getAllUserStmt, err = agendaDB.Prepare(getAllUser)
	checkErr(err)
	addMeetingStmt, err = agendaDB.Prepare(addMeeting)
	checkErr(err)
	getPartStmt, err = agendaDB.Prepare(getParticipating)
	checkErr(err)
	getCreatedStmt, err = agendaDB.Prepare(getCreatedMeeting)
	checkErr(err)
	getTitleMeetingStmt, err = agendaDB.Prepare(getMeetingByTitle)
	checkErr(err)
	addPartStmt, err = agendaDB.Prepare(addParticipant)
	checkErr(err)
	getAllMeetingsStmt, err = agendaDB.Prepare(getAllMeetings)
	checkErr(err)
	getAllParticipantsStmt, err = agendaDB.Prepare(getAllParticipantsOfMeeting)
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
		output = append(output, FindUserInfo{uid: uid, Username: uname, Email: uemail})
	}
	return output, err
}

func AddMeeting(title, creator string, startTime, endTime time.Time) error {
	if startTime.After(endTime) {
		return errors.New("start time is after end time")
	}
	result, err := getUserByNameStmt.Query(creator)
	var uid int
	if err != nil {
		return err
	}
	if result.Next() {
		var name, pass, email, time string
		result.Scan(&uid, &name, &pass, &email, &time)
	} else {
		return errors.New("creator user does not exist when creating meeting")
	}
	result.Close()

	conflict, err := checkUserAvailable(uid, startTime, endTime)
	if err != nil {
		return err
	}
	if conflict != "" {
		return errors.New("time conflict meeting: " + conflict)
	}

	if err = checkMeetingDuplicate(title); err != nil {
		return err
	}

	_, err = addMeetingStmt.Exec(title, uid, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339))
	if err != nil {
		return err
	}
	return nil
}

func checkMeetingDuplicate(title string) error {
	result, err := getTitleMeetingStmt.Query(title)
	if err != nil {
		return err
	}
	defer result.Close()
	if result.Next() {
		return errors.New("duplicate meeting title")
	}
	return nil
}

func checkUserAvailable(uid int, newStart, newEnd time.Time) (conflictMeeting string, err error) {
	result, err := getPartStmt.Query(uid)
	if err != nil {
		return "", err
	}
	if result.Next() {
		var title, start, end string
		result.Scan(&title, &start, &end)
		startTime, _ := time.Parse(time.RFC3339, start)
		endTime, _ := time.Parse(time.RFC3339, end)
		if (newStart.Before(endTime) && newStart.After(startTime)) ||
			(newEnd.Before(endTime) && newEnd.After(startTime)) ||
			(newStart.Before(startTime) && newEnd.After(endTime)) {
				return title, nil
		}
	}
	result.Close()
	result, err = getCreatedStmt.Query(uid)
	defer result.Close()
	if err != nil {
		return "", err
	}
	if result.Next() {
		var title, start, end string
		result.Scan(&title, &start, &end)
		startTime, _ := time.Parse(time.RFC3339, start)
		endTime, _ := time.Parse(time.RFC3339, end)
		if (newStart.Before(endTime) && newStart.After(startTime)) ||
			(newEnd.Before(endTime) && newEnd.After(startTime)) ||
			(newStart.Before(startTime) && newEnd.After(endTime)) {
			return title, nil
		}
	}
	return "", nil
}

func AddPaticipant(title, username string) error {
	result, err := getTitleMeetingStmt.Query(title)
	if err != nil {
		return err
	}
	var mid, creatorID, uid int
	if result.Next() {
		var title, start, end string
		result.Scan(&mid, &title, &creatorID, &start, &end)
	} else {
		return errors.New("meeting does not exist")
	}
	result.Close()
	result, err = getUserByNameStmt.Query(username)
	if err != nil {
		return err
	}
	if result.Next() {
		var name, pass, email, time string
		result.Scan(&uid, &name, &pass, &email, &time)
	} else {
		return errors.New("user does not exist")
	}
	result.Close()
	if uid == creatorID {
		return errors.New("new participant is the creator")
	}
	_, err = addPartStmt.Exec(uid, mid)
	return err
}

// Search meeting by special title
func FindMeetingByTitle(title string) ([]Meeting, error) {
	var (
		output []Meeting
		inFlag bool = false
	)
	curInfo, _ := GetCurrentUser()
	curName := curInfo.Username
	format := "2006-01-02 15:04"
	result, err := getTitleMeetingStmt.Query(title)
	defer result.Close()
	if err != nil {
		return output, err
	}
	if result.Next() {
		var (
			title_, partName, creaName string
			start, end time.Time
			mid, cid, uid int
		)
		result.Scan(&mid, &title_, &cid, &start, &end)
		output = append(output, Meeting{StartTime:start.Format(format), EndTime:end.Format(format), Title:title_})
		result1, err1 := getAllParticipantsStmt.Query(mid)
		defer result1.Close()
		if err1 != nil {
			return output, err1
		}
		// Adding participants of meeting
		for result1.Next() {
			result1.Scan(&uid)
			result2, _ := getUserByIDStmt.Query(uid)
			defer result2.Close()
			if result2.Next() {
				result2.Scan(&partName)
				output[0].Participant = append(output[0].Participant, partName)
				if partName == curName {
					inFlag = true
				}
			}
		}
		// Adding creator of meeting
		result3, _ := getUserByIDStmt.Query(mid)
		result3.Scan(&creaName)
		output[0].Creator = creaName
		if creaName == curName {
			inFlag = true
		}
		result3.Close()
	}
	if inFlag {
		return output, nil
	} else {
		return output, errors.New(fmt.Sprintf("Meeting %v exists, but you are neither its creator nor participant ", title))
	}
}

// search meeting by start time and end time
func FindMeetingsByTimeInterval(start, end time.Time) ([]Meeting, error) {
	var (
		mid, cid, uid int
		startTime, endTime time.Time
		title, creaName, partName string
		cFlag, pFlag bool
		output []Meeting
	)
	curInfo, _ := GetCurrentUser()
	curName := curInfo.Username
	format := "2006-01-02 15:04"
	result, _ := getAllMeetingsStmt.Query()
	defer result.Close()
	for result.Next() {
		cFlag = false
		pFlag = false
		result.Scan(&mid, &title, &cid, &startTime, &endTime)
		// Judging the meeting is within the providing interval or not
		if !(startTime.Equal(start) || startTime.After(start)) && (endTime.Equal(end) || endTime.Before(end)) {
			continue
		}
		// JUdging creator and participant of this meeting
		result1, _ := getUserByIDStmt.Query(mid)
		defer result1.Close()
		result1.Scan(&creaName)
		if curName == creaName {
			cFlag = true
		}
		result2, _ := getAllParticipantsStmt.Query(mid)
		defer result2.Close()
		var partArr []string
		for result2.Next() {
			result2.Scan(&uid)
			result3, _ := getUserByIDStmt.Query(uid)
			defer result3.Close()
			result3.Scan(&partName)
			if partName == curName {
				pFlag = true
			}
			partArr = append(partArr, partName)
		}
		if cFlag || pFlag {
			output = append(output, Meeting{Creator: creaName, Title: title, StartTime: startTime.Format(format), EndTime: endTime.Format(format)})
			for _, ele := range partArr {
				output[0].Participant = append(output[0].Participant, ele)
			}
		}
	}
	return output, nil
}

func Close() {
	agendaDB.Close()
}
