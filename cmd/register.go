package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"syscall"

	"github.com/SYSU532/agenda/entity"
	"github.com/SYSU532/agenda/Log"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
)

var createUserName, createUserPass, createUserEmail string

const emailRegex = "^([A-Za-z0-9]+)@([a-z0-9]+)([.])([a-z]+)$"
const usernameRegex = "^[A-Za-z0-9]+$"
const passwordRegex = "^.{6,}$"

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&createUserName, "username", "u", "", "The username of the new user.")
	registerCmd.Flags().StringVarP(&createUserPass, "password", "p", "", "The password of the new user.")
	registerCmd.Flags().StringVarP(&createUserEmail, "email", "e", "", "The email of the new user.")
}

func checkFormat(origin, regexFormat string) bool {
	format, _ := regexp.Compile(regexFormat)
	return format.MatchString(origin)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Long: fmt.Sprintf(`Register a new user with the input username, password and email.

Usage: %v register -uUserName –password pass –email=a@xxx.com`, os.Args[0]),

	Run: func(cmd *cobra.Command, args []string) {
		// Write init lOG
		Log.WriteLog("Invoke register command to create a new user", 1)
		reader := bufio.NewReader(os.Stdin)
		if createUserName == "" {
			fmt.Print("Enter username: ")
			createUserName, _ = reader.ReadString('\n')
			//trim \n
			createUserName = createUserName[:len(createUserName)-1]
		}
		if createUserPass == "" {
			fmt.Print("Enter password: ")
			bytePass, _ := terminal.ReadPassword(int(syscall.Stdin))
			createUserPass = string(bytePass)
		}
		if createUserEmail == "" {
			fmt.Print("\nEnter Email: ")
			createUserEmail, _ = reader.ReadString('\n')
			createUserEmail = createUserEmail[:len(createUserEmail)-1]
		}
		fmt.Println("\nCreating User...")
		fmt.Printf("Username: %v\n", createUserName)
		fmt.Printf("Password: %v\n", createUserPass)
		fmt.Printf("Email: %v\n", createUserEmail)
		validFormat := true
		if !checkFormat(createUserName, usernameRegex) {
			fmt.Println("Username does not fit the required format!")
			Log.WriteLog("Regist Error: UserName does not fit the required format!", 0)
			validFormat = false
		}
		if !checkFormat(createUserPass, passwordRegex) {
			fmt.Println("Password does not fit the required format!")
			Log.WriteLog("Regist Error: Password does not fit the required format!", 0)
			validFormat = false
		}
		if !checkFormat(createUserEmail, emailRegex) {
			fmt.Println("Email does not fit the required format!")
			Log.WriteLog("Regist Error: Email does not fit the required format!", 0)
			validFormat = false
		}
		if validFormat {
			err := entity.AddUser(createUserName, createUserPass, createUserEmail)
			if err == nil {
				fmt.Println("Successfully created user!")
				Log.WriteLog(fmt.Sprintf("Successfully create user %s, email %s", createUserName, createUserEmail), 1)
				entity.SetCurrentUser(createUserName, createUserPass)
				fmt.Println("Automatically login finished!")
				Log.WriteLog(fmt.Sprintf("Login as user %s succeeded", createUserName), 1)
			} else {
				fmt.Println(err)
				fmt.Println("FAIL to create user!")
				Log.WriteLog(err.Error(), 0)
				Log.WriteLog("FAIL to create user!", 0)
			}
		} else {
			fmt.Println("FAIL to create user!")
			Log.WriteLog("FAIL to create user!", 0)
		}
	},
}
