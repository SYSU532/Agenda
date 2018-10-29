package entity

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"path"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const userInfoPath = "./data"
const userInfoFilename = "curUser.txt"

func GetCurrentUser() (*UserInfo, error) {
	infoFile, err := os.Open(path.Join(userInfoPath, userInfoFilename))
	if err != nil {
		return nil, errors.New("fail to open current user file")
	}
	defer infoFile.Close()
	JsonDecoder := json.NewDecoder(infoFile)

	info := new(UserInfo)
	err = JsonDecoder.Decode(info)
	if err != nil {
		return nil, errors.New("fail to convert Json file to user info")
	}
	pass64, _ := base64.StdEncoding.DecodeString(info.Password)
	info.Password = string(pass64)

	if info.Username == "" || info.Password == "" {
		return nil, errors.New("no current user")
	}
	if err := LoginUser(info.Username, info.Password); err != nil {
		return nil, errors.New("current user info is incorrect: " + err.Error())
	} else {
		return info, nil
	}
}

func SetCurrentUser(username, password string) error {
	infoFile, err := os.Create(path.Join(userInfoPath, userInfoFilename))
	if err != nil {
		return errors.New("fail to open current user file")
	}
	defer infoFile.Close()
	JsonEncoder := json.NewEncoder(infoFile)

	pass64 := base64.StdEncoding.EncodeToString([]byte(password))
	err = JsonEncoder.Encode(&UserInfo{username, pass64})
	if err != nil {
		return errors.New("fail to convert user info to Json")
	}
	return nil
}

func ClearCurrentUser() error {
	infoFile, err := os.Create(path.Join(userInfoPath, userInfoFilename))
	if err != nil {
		return errors.New("fail to open current user file")
	}
	defer infoFile.Close()
	JsonEncoder := json.NewEncoder(infoFile)

	err = JsonEncoder.Encode(struct{}{})
	if err != nil {
		return errors.New("fail to convert user info to Json")
	}
	return nil
}
